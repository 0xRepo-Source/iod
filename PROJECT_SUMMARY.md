# IndexOf Downloader - Project Summary

## What is This?

**IndexOf Downloader** is a command-line tool written in Go that downloads files from Apache-style directory listing websites (like example.com, old archive sites, etc.). It can recursively traverse directories, filter by file patterns, and download files concurrently.

## Key Features

✅ **Recursive Downloads** - Navigate through nested directories  
✅ **Pattern Matching** - Download only specific file types (*.pdf, *.zip, etc.)  
✅ **Concurrent Downloads** - Configurable worker pool for faster downloads  
✅ **Progress Tracking** - Real-time progress bars for each download  
✅ **Smart Filtering** - Filter files by extension or pattern  
✅ **Dry Run Mode** - List files before downloading (list command)  
✅ **Depth Control** - Limit how deep the recursion goes  
✅ **Resume Support** - Skips already downloaded files  

## Project Structure

```
IndexOf/
├── main.go                          # Application entry point
├── go.mod                           # Go module dependencies
├── go.sum                           # Dependency checksums
├── iod.exe           # Compiled binary (Windows)
├── cmd/
│   ├── root.go                      # Root CLI command setup
│   ├── download.go                  # Download command implementation
│   └── list.go                      # List command implementation
├── pkg/
│   └── downloader/
│       └── downloader.go            # Core download logic
├── README.md                        # Full documentation
├── QUICKSTART.md                    # Quick start guide
├── EXAMPLES.md                      # Usage examples
├── Makefile                         # Build automation
└── .gitignore                       # Git ignore rules
```

## Technology Stack

- **Language**: Go 1.21+
- **CLI Framework**: [Cobra](https://github.com/spf13/cobra) - Powerful CLI framework
- **HTML Parsing**: [goquery](https://github.com/PuerkitoBio/goquery) - jQuery-like HTML parsing
- **Progress Bars**: [progressbar](https://github.com/schollz/progressbar) - Beautiful progress indicators

## Quick Commands

```bash
# Build the application
go build -o iod.exe

# List files (no download)
.\iod.exe list https://example.com/pdf/microsoft/ -d 0

# Download files
.\iod.exe download https://example.com/pdf/microsoft/ -o ./downloads

# Download recursively with filters
.\iod.exe download https://example.com/pdf/microsoft/ -r -d 2 -p "*.pdf" -o ./pdfs
```

## How It Works

1. **Fetches** the directory listing HTML from the provided URL
2. **Parses** the HTML to extract file and directory links
3. **Filters** files based on the provided pattern (if any)
4. **Downloads** files using a concurrent worker pool
5. **Tracks** progress with real-time progress bars
6. **Preserves** directory structure in the output folder

## Use Cases

- 📚 Archive historical documentation from sites like example.com
- 💾 Backup files from old FTP-style directory listings
- 📦 Download software packages from archive repositories
- 🔍 Index and catalog files from directory listing sites
- 📥 Batch download files matching specific patterns

## Commands Overview

### `download` Command
Downloads files from a directory listing with support for:
- Output directory specification (`-o`)
- Recursive traversal (`-r`)
- Depth limiting (`-d`)
- Pattern matching (`-p`)
- Concurrent workers (`-w`)
- Verbose output (`-v`)

### `list` Command
Lists files without downloading them, useful for:
- Previewing what's available
- Counting files before download
- Checking file sizes
- Planning download strategy

## Configuration Options

| Option | Flag | Default | Description |
|--------|------|---------|-------------|
| Output Dir | `-o, --output` | `./downloads` | Where to save files |
| Recursive | `-r, --recursive` | `false` | Enable recursive downloads |
| Max Depth | `-d, --depth` | `-1` (unlimited) | Maximum recursion depth |
| Pattern | `-p, --pattern` | `*` (all) | File pattern to match |
| Workers | `-w, --workers` | `3` | Concurrent downloads |
| Verbose | `-v, --verbose` | `false` | Detailed output |

## Example Output

```
$ .\iod.exe download https://example.com/pdf/microsoft/ -d 0 -o ./test

Found 5 files to download

file1.pdf 100% |████████████████████| (1.2 MB/1.2 MB, 2.5 MB/s)
file2.txt 100% |████████████████████| (45 KB/45 KB, 890 KB/s)

✓ Download completed!
  Files downloaded: 5
  Total size: 15.3 MB
  Directories scanned: 1
```

## Performance

- **Concurrent Downloads**: Uses worker pool pattern for efficient downloading
- **Memory Efficient**: Streams files directly to disk
- **Smart Resume**: Skips already downloaded files
- **Rate Limiting**: Configurable workers to respect server limits

## Best Practices

1. **Start with `list`**: Preview files before downloading
2. **Use depth limits**: Avoid downloading entire archives accidentally
3. **Filter by pattern**: Download only what you need
4. **Be respectful**: Use reasonable worker counts (3-5)
5. **Test first**: Try with `-d 0` on a small directory first

## Future Enhancements (Ideas)

- [ ] Resume interrupted downloads
- [ ] Download queue/scheduling
- [ ] Checksum verification
- [ ] Bandwidth limiting
- [ ] Mirror mode (keep in sync)
- [ ] Authentication support
- [ ] Custom headers support
- [ ] Download history/database

## License

MIT License - Free to use and modify!

## Contributing

Contributions welcome! Feel free to:
- Report bugs
- Suggest features
- Submit pull requests
- Improve documentation

## Support

For issues or questions:
1. Check the README.md for detailed documentation
2. Review EXAMPLES.md for usage patterns
3. Read QUICKSTART.md for getting started
4. Open an issue on GitHub (if hosted)

---

**Built with ❤️ in Go**

