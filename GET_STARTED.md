# 🚀 IndexOf Downloader - Universal File Downloader for Directory Listings

## ✨ What You Have

A fully functional Go CLI application that can download **ANY file type** from directory listings:
- ✅ Download **any file type**: PDFs, ZIPs, ISOs, EXEs, images, videos, documents, archives, etc.
- ✅ Works with Apache/Nginx directory listings, FTP archives, example.com, etc.
- ✅ Recursively traverse directories
- ✅ Filter files by pattern (*.pdf, *.zip, *.iso, *.exe, *.mp4, etc.)
- ✅ Download concurrently with multiple workers
- ✅ Show progress bars for each download
- ✅ List files before downloading (dry run)
- ✅ Skip already downloaded files

## 📁 Project Files

```
IndexOf/
├── iod.exe       ← Your compiled binary (ready to use!)
├── main.go                      ← Entry point
├── go.mod                       ← Dependencies
├── go.sum                       ← Dependency checksums
│
├── cmd/                         ← Command implementations
│   ├── root.go                  ← CLI setup
│   ├── download.go              ← Download command
│   └── list.go                  ← List command
│
├── pkg/downloader/              ← Core logic
│   └── downloader.go            ← Download engine
│
└── Documentation/
    ├── README.md                ← Full documentation
    ├── QUICKSTART.md            ← Quick start guide
    ├── EXAMPLES.md              ← Usage examples
    └── PROJECT_SUMMARY.md       ← Project overview
```

## 🎯 Quick Start

### 1. Test It Right Now!

```powershell
# Show help
.\iod.exe --help

# List files from example.com
.\iod.exe list https://example.com/pdf/microsoft/ -d 0

# Download a small test (just current directory, no recursion)
.\iod.exe download https://example.com/pdf/microsoft/ -d 0 -o .\test-download
```

### 2. Real Examples

```powershell
# Download all PDF files recursively (depth 2) from Microsoft docs
.\iod.exe download https://example.com/pdf/microsoft/ -r -d 2 -p "*.pdf" -o .\microsoft-pdfs

# List all files recursively with verbose output
.\iod.exe list https://example.com/pdf/microsoft/ -r -v

# Download with 5 concurrent workers
.\iod.exe download https://example.com/pdf/microsoft/ -r -w 5 -o .\downloads
```

## 🔧 Commands

### `download` - Download Files
```
.\iod.exe download <URL> [flags]

Flags:
  -o, --output DIR       Output directory (default: ./downloads)
  -r, --recursive        Download recursively
  -d, --depth N          Max depth (-1 = unlimited)
  -p, --pattern GLOB     File pattern (e.g., *.pdf)
  -w, --workers N        Concurrent downloads (default: 3)
  -v, --verbose          Verbose output
```

### `list` - List Files (No Download)
```
.\iod.exe list <URL> [flags]

Flags:
  -r, --recursive        List recursively
  -d, --depth N          Max depth (-1 = unlimited)
  -p, --pattern GLOB     File pattern (e.g., *.pdf)
  -v, --verbose          Show file sizes
```

## 💡 Common Use Cases

### 1. Archive Historical Documentation
```powershell
.\iod.exe download https://example.com/pdf/microsoft/ -r -d 3 -p "*.pdf" -o .\archive
```

### 2. Preview Before Downloading
```powershell
# First, see what's there
.\iod.exe list https://example.com/pdf/microsoft/basic/ -r -v

# Then download
.\iod.exe download https://example.com/pdf/microsoft/basic/ -r -o .\basic-docs
```

### 3. Download Specific File Types
```powershell
# Only PDFs
.\iod.exe download URL -r -p "*.pdf" -o .\pdfs

# Only ZIP files
.\iod.exe download URL -r -p "*.zip" -o .\archives

# Only text files
.\iod.exe download URL -r -p "*.txt" -o .\text-files
```

## 🛠️ Development

### Build from Source
```powershell
# Build for Windows
go build -o iod.exe

# Build for Linux
$env:GOOS="linux"; go build -o iod

# Build for macOS
$env:GOOS="darwin"; go build -o iod
```

### Run Without Building
```powershell
go run . list https://example.com/pdf/microsoft/ -d 0
go run . download URL -o ./test
```

### Install Globally
```powershell
go install .
# Then use from anywhere as: iod
```

## 📊 Example Output

```
PS G:\IndexOf> .\iod.exe download https://example.com/pdf/microsoft/ -d 0

Found 5 files to download

file1.pdf 100% |████████████████████| (1.2 MB/1.2 MB)
file2.txt 100% |████████████████████| (45 KB/45 KB)
file3.pdf 100% |████████████████████| (856 KB/856 KB)

✓ Download completed!
  Files downloaded: 5
  Total size: 15.3 MB
  Directories scanned: 1
```

## 🎓 Tips

1. **Always test with `list` first** - See what's available before downloading
2. **Use depth limits** - Start with `-d 1` or `-d 2` to avoid downloading too much
3. **Be respectful** - Use 3-5 workers max; don't overwhelm servers
4. **Filter by pattern** - Use `-p "*.pdf"` to download only what you need
5. **Check output directory** - Make sure you have enough disk space
6. **Use verbose mode** - Add `-v` to see detailed progress

## 🔍 Troubleshooting

### Build Issues
```powershell
# Clean and rebuild
go clean
go mod tidy
go build -o iod.exe
```

### Runtime Issues
```powershell
# Check if the binary exists
Get-Item .\iod.exe

# Run with verbose output to see what's happening
.\iod.exe download URL -v
```

### Download Failures
- Check your internet connection
- Verify the URL is correct and accessible
- Try reducing the number of workers (`-w 1`)
- Some servers may rate-limit; be patient

## 📚 Documentation

- **README.md** - Complete documentation
- **QUICKSTART.md** - Getting started guide
- **EXAMPLES.md** - Real-world examples
- **PROJECT_SUMMARY.md** - Technical overview

## 🚀 Next Steps

1. **Try it out** with a small test:
   ```powershell
   .\iod.exe list https://example.com/pdf/microsoft/ -d 0
   ```

2. **Download something small**:
   ```powershell
   .\iod.exe download https://example.com/pdf/microsoft/ -d 0 -o .\test
   ```

3. **Explore the codebase** in `cmd/` and `pkg/downloader/`

4. **Customize** the tool for your specific needs

5. **Share** with others who might find it useful!

## 🌟 Features Implemented

- [x] CLI with Cobra framework
- [x] HTML parsing with goquery
- [x] Concurrent downloads with worker pool
- [x] Progress bars
- [x] Pattern matching (glob patterns)
- [x] Recursive directory traversal
- [x] Depth limiting
- [x] File skipping (already downloaded)
- [x] Verbose mode
- [x] List command (dry run)
- [x] Download statistics
- [x] Error handling

## 📝 License

MIT License - Free to use and modify!

---

**You're all set! The application is ready to use. Try the quick start commands above to get started!** 🎉

