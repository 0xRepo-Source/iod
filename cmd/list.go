package cmd

import (
	"fmt"

	"github.com/iod/pkg/downloader"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [URL]",
	Short: "List files from a directory listing without downloading",
	Long:  `List all files from a directory listing website without downloading them`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]

		// Announce the bot
		fmt.Println("🤖 Archive Bot v1.0.0 - Scanning directory structure...")
		fmt.Printf("📍 Target: %s\n", url)
		if recursive {
			fmt.Println("🔄 Mode: Recursive scan")
		} else {
			fmt.Println("📄 Mode: Single directory")
		}
		fmt.Println()

		config := downloader.Config{
			BaseURL:     url,
			Recursive:   recursive,
			MaxDepth:    maxDepth,
			FilePattern: filePattern,
			Verbose:     verbose,
		}

		dl := downloader.NewDownloader(config)
		files, err := dl.ListFiles()
		if err != nil {
			return fmt.Errorf("failed to list files: %w", err)
		}

		fmt.Printf("📋 Archive Bot found %d files:\n\n", len(files))
		for _, file := range files {
			if verbose {
				fmt.Printf("  %s (%s)\n", file.URL, formatBytes(file.Size))
			} else {
				fmt.Printf("  %s\n", file.URL)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "List files recursively")
	listCmd.Flags().IntVarP(&maxDepth, "depth", "d", -1, "Maximum depth for recursive listing (-1 for unlimited)")
	listCmd.Flags().StringVarP(&filePattern, "pattern", "p", "*", "File pattern to match (e.g., *.pdf, *.zip)")
}
