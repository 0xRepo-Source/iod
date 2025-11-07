package downloader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/schollz/progressbar/v3"
)

type Config struct {
	BaseURL     string
	OutputDir   string
	Recursive   bool
	MaxDepth    int
	FilePattern string
	Workers     int
	Verbose     bool
}

type Downloader struct {
	config     Config
	httpClient *http.Client
	stats      *DownloadStats
	mu         sync.Mutex
}

type DownloadStats struct {
	FilesDownloaded int
	TotalBytes      int64
	DirsScanned     int
	FailedDownloads int
}

type FileInfo struct {
	URL  string
	Path string
	Size int64
}

func NewDownloader(config Config) *Downloader {
	if config.Workers <= 0 {
		config.Workers = 3
	}

	return &Downloader{
		config: config,
		httpClient: &http.Client{
			Timeout: 30 * time.Minute,
		},
		stats: &DownloadStats{},
	}
}

func (d *Downloader) Download() (*DownloadStats, error) {
	files, err := d.ListFiles()
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		fmt.Println("No files found to download")
		return d.stats, nil
	}

	fmt.Printf("Found %d files to download\n\n", len(files))

	// Create a work queue
	jobs := make(chan FileInfo, len(files))
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < d.config.Workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for file := range jobs {
				if err := d.downloadFile(file); err != nil {
					d.mu.Lock()
					d.stats.FailedDownloads++
					d.mu.Unlock()
					if d.config.Verbose {
						fmt.Printf("✗ Failed to download %s: %v\n", file.URL, err)
					}
				} else {
					d.mu.Lock()
					d.stats.FilesDownloaded++
					d.mu.Unlock()
				}
			}
		}(i)
	}

	// Queue all files
	for _, file := range files {
		jobs <- file
	}
	close(jobs)

	// Wait for all downloads to complete
	wg.Wait()

	return d.stats, nil
}

func (d *Downloader) ListFiles() ([]FileInfo, error) {
	var files []FileInfo
	visited := make(map[string]bool)

	err := d.scanDirectory(d.config.BaseURL, "", 0, visited, &files)
	if err != nil {
		return nil, err
	}

	// Filter files by pattern
	if d.config.FilePattern != "" && d.config.FilePattern != "*" {
		files = d.filterByPattern(files)
	}

	return files, nil
}

func (d *Downloader) scanDirectory(dirURL, relativePath string, depth int, visited map[string]bool, files *[]FileInfo) error {
	// Check if we've already visited this URL
	if visited[dirURL] {
		return nil
	}
	visited[dirURL] = true

	// Check depth limit
	if d.config.MaxDepth >= 0 && depth > d.config.MaxDepth {
		return nil
	}

	d.mu.Lock()
	d.stats.DirsScanned++
	d.mu.Unlock()

	if d.config.Verbose {
		fmt.Printf("Scanning: %s\n", dirURL)
	}

	// Fetch the directory listing
	resp, err := d.httpClient.Get(dirURL)
	if err != nil {
		return fmt.Errorf("failed to fetch %s: %w", dirURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d for %s", resp.StatusCode, dirURL)
	}

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Extract links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		// Skip parent directory and special links
		if href == "../" || href == "/" || strings.HasPrefix(href, "?") || strings.HasPrefix(href, "#") {
			return
		}

		// Skip absolute URLs to other domains or protocol-relative URLs
		if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") || strings.HasPrefix(href, "//") {
			baseURL, _ := url.Parse(dirURL)

			// Skip if it's not under our base path
			if !strings.HasPrefix(href, baseURL.Scheme+"://"+baseURL.Host+baseURL.Path) {
				return
			}
		}

		// Parse the URL
		linkURL, err := url.Parse(href)
		if err != nil {
			return
		}

		// Resolve relative URLs
		baseURL, _ := url.Parse(dirURL)
		absoluteURL := baseURL.ResolveReference(linkURL).String()

		// Skip if the resolved URL goes outside our base directory
		if !strings.HasPrefix(absoluteURL, d.config.BaseURL) {
			return
		}

		// Determine if it's a directory (ends with /)
		isDir := strings.HasSuffix(href, "/")

		if isDir {
			// Recursively scan subdirectory
			if d.config.Recursive {
				newRelativePath := filepath.Join(relativePath, strings.TrimSuffix(href, "/"))
				d.scanDirectory(absoluteURL, newRelativePath, depth+1, visited, files)
			}
		} else {
			// It's a file
			filePath := filepath.Join(relativePath, filepath.Base(href))
			*files = append(*files, FileInfo{
				URL:  absoluteURL,
				Path: filePath,
			})
		}
	})

	return nil
}

func (d *Downloader) downloadFile(file FileInfo) error {
	// Create output path
	outputPath := filepath.Join(d.config.OutputDir, file.Path)

	// Create directory structure
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Check if file already exists
	if _, err := os.Stat(outputPath); err == nil {
		if d.config.Verbose {
			fmt.Printf("⊘ Skipping (already exists): %s\n", file.Path)
		}
		return nil
	}

	// Download the file
	resp, err := d.httpClient.Get(file.URL)
	if err != nil {
		return fmt.Errorf("failed to download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Create output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer outFile.Close()

	// Create progress bar
	var reader io.Reader = resp.Body
	if resp.ContentLength > 0 {
		bar := progressbar.DefaultBytes(
			resp.ContentLength,
			filepath.Base(file.Path),
		)
		reader = io.TeeReader(resp.Body, bar)
	}

	// Copy data
	written, err := io.Copy(outFile, reader)
	if err != nil {
		os.Remove(outputPath) // Clean up partial download
		return fmt.Errorf("failed to write file: %w", err)
	}

	d.mu.Lock()
	d.stats.TotalBytes += written
	d.mu.Unlock()

	if d.config.Verbose {
		fmt.Printf("✓ Downloaded: %s\n", file.Path)
	}

	return nil
}

func (d *Downloader) filterByPattern(files []FileInfo) []FileInfo {
	var filtered []FileInfo
	pattern := d.config.FilePattern

	for _, file := range files {
		matched, err := filepath.Match(pattern, filepath.Base(file.Path))
		if err == nil && matched {
			filtered = append(filtered, file)
		}
	}

	return filtered
}
