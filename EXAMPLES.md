# Usage Examples

## Basic Examples

### 1. List Files from a Directory

```bash
# List files in the Microsoft directory (depth 0 - current directory only)
.\iod.exe list https://example.com/pdf/microsoft/ -d 0

# List all files recursively
.\iod.exe list https://example.com/pdf/microsoft/ -r

# List only PDF files
.\iod.exe list https://example.com/pdf/microsoft/ -r -p "*.pdf"
```

### 2. Download Files

```bash
# Download files from current directory only
.\iod.exe download https://example.com/pdf/microsoft/ -o ./microsoft-docs

# Download recursively with depth limit
.\iod.exe download https://example.com/pdf/microsoft/ -r -d 2 -o ./microsoft-docs

# Download only specific file types
.\iod.exe download https://example.com/pdf/microsoft/ -r -p "*.pdf" -o ./pdfs

# Download with more concurrent workers for faster downloads
.\iod.exe download https://example.com/pdf/microsoft/ -r -w 5 -o ./downloads
```

### 3. Advanced Usage

```bash
# Download only text files, verbose output
.\iod.exe download https://example.com/pdf/microsoft/ -p "*.txt" -o ./text-files -v

# List all ZIP files recursively with size information
.\iod.exe list https://example.com/pdf/microsoft/ -r -p "*.zip" -v

# Download from a specific subdirectory
.\iod.exe download https://example.com/pdf/microsoft/basic/ -r -o ./basic-docs
```

## Real-World Scenarios

### Scenario 1: Archive Historical Documentation

```bash
# Download all PDFs from Microsoft documentation archive
.\iod.exe download https://example.com/pdf/microsoft/ -r -d 3 -p "*.pdf" -o ./microsoft-archive -w 5
```

### Scenario 2: Preview Before Downloading

```bash
# First, see what's available
.\iod.exe list https://example.com/pdf/microsoft/languages/ -r -v

# Then download specific files
.\iod.exe download https://example.com/pdf/microsoft/languages/ -r -p "*.pdf" -o ./languages
```

### Scenario 3: Download from Multiple Sources

```bash
# Download from different directories
.\iod.exe download https://example.com/pdf/microsoft/basic/ -r -o ./downloads/basic
.\iod.exe download https://example.com/pdf/microsoft/c/ -r -o ./downloads/c
.\iod.exe download https://example.com/pdf/microsoft/fortran/ -r -o ./downloads/fortran
```

## Tips

1. **Use depth limits** (`-d`) to avoid downloading too much at once
2. **Start with list command** to see what's available before downloading
3. **Use patterns** (`-p`) to download only specific file types
4. **Adjust workers** (`-w`) based on your internet speed and server limits
5. **Enable verbose mode** (`-v`) to see detailed progress information
6. **Check the output directory** before starting large downloads

## Pattern Matching Examples

```bash
# Download only PDF files
-p "*.pdf"

# Download only ZIP archives
-p "*.zip"

# Download only text files
-p "*.txt"

# Download only documentation (txt and pdf)
# Note: Use multiple commands for OR patterns
.\iod.exe download URL -p "*.pdf" -o ./downloads
.\iod.exe download URL -p "*.txt" -o ./downloads
```

## Performance Tuning

```bash
# Fast download with 10 workers (be respectful to servers!)
.\iod.exe download URL -r -w 10 -o ./downloads

# Slower, more conservative download with 2 workers
.\iod.exe download URL -r -w 2 -o ./downloads

# Default balanced approach (3 workers)
.\iod.exe download URL -r -o ./downloads
```

