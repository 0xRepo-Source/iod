# Changelog

All notable changes to this project will be documented in this file.

## [1.0.0] - 2025-11-07

### Added
- Initial release of IOD (IndexOf Downloader)
- Universal file downloader for directory listings
- Recursive directory traversal with depth control
- Pattern matching for file filtering
- Concurrent downloads with configurable workers
- Progress bars for downloads
- List command for dry-run preview
- Resume support (skips existing files)
- Comprehensive documentation
- Support for Apache/Nginx directory listings
- Works with all file types (PDFs, ZIPs, ISOs, videos, etc.)

### Features
- `download` command - Download files from directory listings
- `list` command - Preview files without downloading
- `-r, --recursive` - Enable recursive downloads
- `-d, --depth` - Control recursion depth
- `-p, --pattern` - Filter files by glob pattern
- `-w, --workers` - Configure concurrent downloads
- `-v, --verbose` - Enable detailed output
- `-o, --output` - Specify output directory

### Documentation
- README.md - Main documentation
- QUICKSTART.md - Quick start guide
- RECURSIVE_DOWNLOAD.md - Full recursive download guide
- FILE_TYPES.md - Complete file type examples
- HOW_RECURSION_WORKS.md - Visual recursion explanation
- EXAMPLES.md - Real-world usage examples
- GET_STARTED.md - Getting started guide
- CONTRIBUTING.md - Contribution guidelines

[1.0.0]: https://github.com/yourusername/iod/releases/tag/v1.0.0
