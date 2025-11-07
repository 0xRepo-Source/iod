# Quick Start Guide

## Installation

1. Make sure you have Go 1.21+ installed:
   ```bash
   go version
   ```

2. Navigate to the project directory:
   ```bash
   cd g:\IndexOf
   ```

3. Download dependencies:
   ```bash
   go mod download
   ```

4. Build the application:
   ```bash
   go build -o iod.exe
   ```

## First Steps

### 1. Test the Application

```bash
# Show help
.\iod.exe --help

# List available commands
.\iod.exe
```

### 2. List Files (No Download)

```bash
# List files from example.com Microsoft directory
.\iod.exe list https://example.com/pdf/microsoft/ -d 0
```

### 3. Download a Small Set First

```bash
# Download just the current directory (no recursion)
.\iod.exe download https://example.com/pdf/microsoft/ -d 0 -o ./test-download
```

### 4. Download Recursively

```bash
# Download with depth limit and pattern matching
.\iod.exe download https://example.com/pdf/microsoft/ -r -d 2 -p "*.pdf" -o ./downloads
```

## Common Commands Cheat Sheet

| Command | Description |
|---------|-------------|
| `.\iod.exe --help` | Show general help |
| `.\iod.exe download --help` | Show download command help |
| `.\iod.exe list URL` | List files from URL |
| `.\iod.exe list URL -r` | List files recursively |
| `.\iod.exe download URL -o DIR` | Download to DIR |
| `.\iod.exe download URL -r -d 3` | Download recursively, max depth 3 |
| `.\iod.exe download URL -p "*.pdf"` | Download only PDF files |
| `.\iod.exe download URL -w 5` | Use 5 concurrent workers |
| `.\iod.exe download URL -v` | Verbose output |

## Flags Reference

### Global Flags
- `-v, --verbose` - Enable verbose output

### Download Command Flags
- `-o, --output DIR` - Output directory (default: `./downloads`)
- `-r, --recursive` - Download recursively
- `-d, --depth N` - Max recursion depth (-1 = unlimited)
- `-p, --pattern PATTERN` - File pattern (e.g., `*.pdf`, `*.zip`)
- `-w, --workers N` - Number of concurrent downloads (default: 3)

### List Command Flags
- `-r, --recursive` - List recursively
- `-d, --depth N` - Max recursion depth (-1 = unlimited)
- `-p, --pattern PATTERN` - File pattern to match

## Tips for Beginners

1. **Always test with list first**: Use the `list` command to see what files exist before downloading
2. **Start small**: Use `-d 0` or `-d 1` to limit depth when testing
3. **Be respectful**: Don't use too many workers (`-w`) as it may overwhelm servers
4. **Use patterns**: Filter by file type with `-p "*.pdf"` to avoid downloading unwanted files
5. **Check disk space**: Large archives can be several GB or more

## Troubleshooting

### Application won't run
```bash
# Make sure it's built
go build -o iod.exe

# Check if the executable exists
Get-Item .\iod.exe
```

### Dependencies not found
```bash
# Download dependencies
go mod download
go mod tidy
```

### Build errors
```bash
# Clean and rebuild
go clean
go build -o iod.exe
```

## Example Workflow

```bash
# 1. First, explore what's available
.\iod.exe list https://example.com/pdf/microsoft/ -r -d 1

# 2. Check how many PDFs there are
.\iod.exe list https://example.com/pdf/microsoft/ -r -p "*.pdf" -v

# 3. Download just the PDFs with reasonable depth
.\iod.exe download https://example.com/pdf/microsoft/ -r -d 2 -p "*.pdf" -o ./microsoft-pdfs

# 4. Check what was downloaded
Get-ChildItem -Recurse .\microsoft-pdfs
```

## Building for Other Platforms

```bash
# For Linux
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o iod-linux

# For macOS (Intel)
$env:GOOS="darwin"; $env:GOARCH="amd64"; go build -o iod-macos

# For macOS (Apple Silicon)
$env:GOOS="darwin"; $env:GOARCH="arm64"; go build -o iod-macos-arm

# Reset to Windows
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o iod.exe
```

## Next Steps

- Read the full [README.md](README.md) for detailed documentation
- Check [EXAMPLES.md](EXAMPLES.md) for more usage examples
- Star the repository if you find it useful!
- Report issues or contribute improvements

