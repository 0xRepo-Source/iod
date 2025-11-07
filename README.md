# IndexOf Downloader

A **universal file downloader** CLI tool written in Go for downloading **ANY file type** from directory listing websites like example.com, old FTP archives, and Apache-style index pages.

> **📖 Want to download entire directory trees?** See [RECURSIVE_DOWNLOAD.md](RECURSIVE_DOWNLOAD.md) for a complete guide on full recursive downloads.
> 
> **📦 Works with all file types!** See [FILE_TYPES.md](FILE_TYPES.md) for examples of PDFs, ZIPs, ISOs, videos, images, executables, and more.

## Features

- 📥 **Download ANY file type** - PDFs, ZIPs, ISOs, EXEs, images, videos, documents, etc.
- 🔄 **Full recursive directory traversal** - Download entire directory trees
- 🎯 File pattern matching (e.g., `*.pdf`, `*.zip`, `*.iso`, `*.exe`)
- ⚡ Concurrent downloads with configurable workers
- 📊 Progress bars for downloads
- 📝 List files without downloading
- 🎨 Colorful and informative output
- 🚀 Fast and efficient
- 🔁 Resume support - Skips already downloaded files
- 🌐 Works with Apache/Nginx directory listings, old FTP sites, archive repositories

## Installation

### Prerequisites

- Go 1.21 or higher

### Build from Source

```bash
# Clone or navigate to the repository
cd IndexOf

# Download dependencies
go mod download

# Build the binary
go build -o iod.exe

# Or install globally
go install
```

## Usage

### Download Files

Download all files from a directory:
```bash
iod download https://example.com/pdf/microsoft/ -o ./downloads
```

**Download everything recursively (all subdirectories):**
```bash
iod download https://example.com/pdf/microsoft/ -r -o ./downloads
```

Download recursively with a maximum depth of 3:
```bash
iod download https://example.com/pdf/microsoft/ -r -d 3 -o ./downloads
```

Download only specific file types:
```bash
# Only PDFs
iod download https://example.com/pdf/microsoft/ -r -p "*.pdf" -o ./downloads

# Only ZIP archives
iod download https://example.com/archive/ -r -p "*.zip" -o ./archives

# Only ISO images
iod download https://example.com/isos/ -r -p "*.iso" -o ./isos

# Only executables
iod download https://example.com/software/ -r -p "*.exe" -o ./software
```

Download with 5 concurrent workers:
```bash
iod download https://example.com/pdf/microsoft/ -r -w 5 -o ./downloads
```

### List Files

List all files without downloading:
```bash
iod list https://example.com/pdf/microsoft/
```

List files recursively with verbose output:
```bash
iod list https://example.com/pdf/microsoft/ -r -v
```

List only ZIP files:
```bash
iod list https://example.com/pdf/microsoft/ -r -p "*.zip"
```

## Command Reference

### Global Flags

- `-v, --verbose`: Enable verbose output

### Download Command

```
iod download [URL] [flags]
```

**Flags:**
- `-o, --output string`: Output directory for downloaded files (default: `./downloads`)
- `-r, --recursive`: Download recursively
- `-d, --depth int`: Maximum depth for recursive download (-1 for unlimited, default: -1)
- `-p, --pattern string`: File pattern to match, e.g., `*.pdf`, `*.zip` (default: `*`)
- `-w, --workers int`: Number of concurrent downloads (default: 3)

### List Command

```
iod list [URL] [flags]
```

**Flags:**
- `-r, --recursive`: List files recursively
- `-d, --depth int`: Maximum depth for recursive listing (-1 for unlimited, default: -1)
- `-p, --pattern string`: File pattern to match, e.g., `*.pdf`, `*.zip` (default: `*`)

## Examples

### Download Microsoft PDFs from example.com

```bash
# Download all files from the microsoft directory
iod download https://example.com/pdf/microsoft/ -o ./microsoft-docs

# Download only PDFs recursively (depth 2)
iod download https://example.com/pdf/microsoft/ -r -d 2 -p "*.pdf" -o ./microsoft-pdfs
```

### List Files Before Downloading

```bash
# First, see what files are available
iod list https://example.com/pdf/microsoft/ -r -v

# Then download specific patterns
iod download https://example.com/pdf/microsoft/ -r -p "*.pdf" -o ./downloads
```

## Output

The tool provides:
- Real-time progress bars for each download
- Summary statistics including:
  - Number of files downloaded
  - Total size downloaded
  - Number of directories scanned
  - Number of failed downloads (if any)

## How It Works

1. **Parsing**: The tool fetches the directory listing HTML and parses it using goquery
2. **Discovery**: It extracts all file and directory links from the page
3. **Filtering**: Applies your file pattern filter (if specified)
4. **Downloading**: Downloads files concurrently using a worker pool
5. **Organization**: Preserves the directory structure in the output folder

## Dependencies

- [cobra](https://github.com/spf13/cobra) - CLI framework
- [goquery](https://github.com/PuerkitoBio/goquery) - HTML parsing
- [progressbar](https://github.com/schollz/progressbar) - Progress bars

## License

MIT License - feel free to use this tool for any purpose!

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

