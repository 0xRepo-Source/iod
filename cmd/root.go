package cmd

import (
	"github.com/spf13/cobra"
)

var (
	outputDir   string
	recursive   bool
	maxDepth    int
	filePattern string
	workers     int
	verbose     bool
)

var rootCmd = &cobra.Command{
	Use:   "iod",
	Short: "Archive Bot - Download files from directory listing websites",
	Long: `🤖 Archive Bot - Automated archival tool for directory listing websites

A CLI tool to download and archive files from directory listing websites
like bitsavers.org, archive.org, and other index pages. Recursively crawls
and preserves entire directory structures.`,
	Example: `  # Download all files from a directory
  iod download https://bitsavers.org/pdf/microsoft/ -o ./downloads

  # Download recursively with max depth of 3
  iod download https://bitsavers.org/pdf/microsoft/ -r -d 3 -o ./downloads

  # Download only PDF files
  iod download https://bitsavers.org/pdf/microsoft/ -p "*.pdf" -o ./downloads

  # List files without downloading
  iod list https://bitsavers.org/pdf/microsoft/ -r`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}
