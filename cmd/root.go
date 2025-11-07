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
	Short: "Download files from directory listing websites",
	Long: `A CLI tool to download and index files from directory listing websites
like example.com. It can recursively download files matching specific patterns.`,
	Example: `  # Download all files from a directory
  iod download https://example.com/pdf/microsoft/ -o ./downloads

  # Download recursively with max depth of 3
  iod download https://example.com/pdf/microsoft/ -r -d 3 -o ./downloads

  # Download only PDF files
  iod download https://example.com/pdf/microsoft/ -p "*.pdf" -o ./downloads

  # List files without downloading
  iod list https://example.com/pdf/microsoft/ -r`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}
