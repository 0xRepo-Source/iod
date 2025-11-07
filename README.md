# IndexOf Downloader (iod)

A CLI tool for downloading files from directory listing websites (Apache/Nginx directory listings, old FTP archives, bitsavers.org, etc.).

Uses "Archive Bot" user-agent to properly identify itself to web servers.

## Features

- 📥 **Universal file downloader** - Downloads ANY file type (PDFs, ZIPs, ISOs, EXEs, videos, images, documents, ROMs, archives, etc.)
- 🔄 **Recursive directory traversal** - Download entire directory trees with unlimited depth
- 🎯 **Pattern matching** - Filter files by extension (e.g., `*.pdf`, `*.zip`, `*.iso`, `*.exe`)
- ⚡ **Concurrent downloads** - Configurable worker pool for faster downloads
- 📊 **Progress tracking** - Real-time progress bars for each file
- 📝 **Preview mode** - List files without downloading
-  **Resume support** - Automatically skips already downloaded files
- 🌐 **Universal compatibility** - Works with Apache/Nginx listings, FTP archives, bitsavers.org, and more

## Installation

### Prerequisites

- Go 1.21 or higher

### Build from Source

```bash
# Clone the repository
git clone https://github.com/0xRepo-Source/iod.git
cd iod

# Download dependencies
go mod download

# Build the binary
go build -o iod.exe

# Or install globally
go install
```

## Quick Start

```bash
# Download all files from a URL
iod download https://example.com -o ./downloads

# Download recursively (all subdirectories)
iod download https://example.com -r -o ./downloads

# Download only PDFs, recursively
iod download https://example.com -r -p "*.pdf" -o ./pdfs

# Preview files before downloading
iod list https://example.com -r -v
```

## Usage

### Download Command

**Basic download:**
```bash
iod download https://example.com -o ./downloads
```

**Recursive download (all subdirectories):**
```bash
iod download https://example.com -r -o ./downloads
```

**Control recursion depth:**
```bash
iod download https://example.com -r -d 3 -o ./downloads
```

**Filter by file type:**
```bash
# Only PDFs
iod download https://example.com -r -p "*.pdf" -o ./downloads

# Only ZIP files
iod download https://archive.org/download/something/ -r -p "*.zip" -o ./archives

# Only ISO images
iod download https://example.org/isos/ -r -p "*.iso" -o ./isos
```

**Faster downloads with multiple workers:**
```bash
iod download https://example.com -r -w 10 -o ./downloads
```

### List Command

**Preview files without downloading:**
```bash
iod list https://example.com
```

**List recursively with details:**
```bash
iod list https://example.com -r -v
```

**List specific file types:**
```bash
iod list https://example.com -r -p "*.pdf"
```

## Command Reference

### Global Flags

- `-v, --verbose`: Enable verbose output

### Download Command

```bash
iod download [URL] [flags]
```

**Flags:**
- `-o, --output string`: Output directory (default: `./downloads`)
- `-r, --recursive`: Download recursively through subdirectories
- `-d, --depth int`: Maximum recursion depth, -1 for unlimited (default: -1)
- `-p, --pattern string`: File pattern filter (e.g., `*.pdf`, `*.zip`) (default: `*`)
- `-w, --workers int`: Number of concurrent downloads (default: 3)

### List Command

```bash
iod list [URL] [flags]
```

**Flags:**
- `-r, --recursive`: List files recursively
- `-d, --depth int`: Maximum depth for listing (default: -1)
- `-p, --pattern string`: File pattern filter (default: `*`)

## Real-World Examples

### Archive a Website's PDF Collection

```bash
# Download all PDFs from bitsavers.org
iod download https://example.com -r -p "*.pdf" -o ./bitsavers-microsoft
```

### Download Software Archives

```bash
# Download all ZIP files
iod download https://archive.org/download/software_collection/ -r -p "*.zip" -o ./software

# Download all ISO images
iod download https://example.org/isos/linux/ -r -p "*.iso" -o ./isos
```

### Fast Batch Download

```bash
# Use 10 workers for faster downloading
iod download https://bitsavers.org/pdf/dec/ -r -w 10 -o ./dec-docs
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

