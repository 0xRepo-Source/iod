package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iod/pkg/downloader"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download [URL]",
	Short: "Download files from a directory listing",
	Long:  `Download files from a directory listing website like example.com`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]

		// Announce the bot
		fmt.Println("🤖 Archive Bot v1.0.0 - Initiating download sequence...")
		fmt.Printf("📍 Target: %s\n", url)
		if recursive {
			fmt.Println("🔄 Mode: Recursive archival")
		} else {
			fmt.Println("📄 Mode: Single directory")
		}
		fmt.Println()

		// Create output directory if it doesn't exist
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Resolve absolute path
		absOutputDir, err := filepath.Abs(outputDir)
		if err != nil {
			return fmt.Errorf("failed to resolve output directory: %w", err)
		}

		config := downloader.Config{
			BaseURL:     url,
			OutputDir:   absOutputDir,
			Recursive:   recursive,
			MaxDepth:    maxDepth,
			FilePattern: filePattern,
			Workers:     workers,
			Verbose:     verbose,
		}

		dl := downloader.NewDownloader(config)
		stats, err := dl.Download()
		if err != nil {
			return fmt.Errorf("download failed: %w", err)
		}

		fmt.Printf("\n✅ Archive Bot completed successfully!\n")
		fmt.Printf("  📦 Files archived: %d\n", stats.FilesDownloaded)
		fmt.Printf("  💾 Total size: %s\n", formatBytes(stats.TotalBytes))
		fmt.Printf("  📂 Directories scanned: %d\n", stats.DirsScanned)
		if stats.FailedDownloads > 0 {
			fmt.Printf("  ⚠️  Failed downloads: %d\n", stats.FailedDownloads)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&outputDir, "output", "o", "./downloads", "Output directory for downloaded files")
	downloadCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Download recursively")
	downloadCmd.Flags().IntVarP(&maxDepth, "depth", "d", -1, "Maximum depth for recursive download (-1 for unlimited)")
	downloadCmd.Flags().StringVarP(&filePattern, "pattern", "p", "*", "File pattern to match (e.g., *.pdf, *.zip)")
	downloadCmd.Flags().IntVarP(&workers, "workers", "w", 3, "Number of concurrent downloads")
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
