# Full Recursive Download Guide

## Universal Downloader for ANY File Type

The IndexOf Downloader works with **ALL file types** - PDFs, ZIPs, ISOs, EXEs, videos, images, documents, archives, and more!

## How to Download Everything from a Directory (Including All Subdirectories)

The IndexOf Downloader supports full recursive downloads with improved path filtering.

## Basic Full Download

To download **all files** from a directory and **all its subdirectories**:

```powershell
.\iod.exe download https://example.com/pdf/microsoft/ -r -o .\downloads
```

**Flags explained:**
- `-r` = Recursive (download from subdirectories too)
- `-o .\downloads` = Output to the "downloads" folder
- No `-d` flag = Unlimited depth (downloads everything)

## Controlled Full Download

### With Depth Limit

Download everything but limit how deep it goes:

```powershell
# Download 2 levels deep
.\iod.exe download https://example.com/pdf/microsoft/ -r -d 2 -o .\downloads

# Download 3 levels deep
.\iod.exe download https://example.com/pdf/microsoft/ -r -d 3 -o .\downloads
```

**What the depth levels mean:**
- `-d 0` = Current directory only (no subdirectories)
- `-d 1` = Current directory + 1 level of subdirectories
- `-d 2` = Current directory + 2 levels of subdirectories
- `-d -1` or no flag = Unlimited depth (all subdirectories)

### With File Type Filter (Works with ANY file extension!)

Download all files of a specific type, recursively:

```powershell
# Download all PDFs from all subdirectories
.\iod.exe download https://example.com/pdf/microsoft/ -r -p "*.pdf" -o .\pdfs

# Download all ZIP/RAR archives from all subdirectories
.\iod.exe download https://example.com/archives/ -r -p "*.zip" -o .\archives
.\iod.exe download https://example.com/archives/ -r -p "*.rar" -o .\archives

# Download all ISO images
.\iod.exe download https://example.com/isos/ -r -p "*.iso" -o .\isos

# Download all executables
.\iod.exe download https://example.com/software/ -r -p "*.exe" -o .\software

# Download all videos (MP4, AVI, MKV, etc.)
.\iod.exe download https://example.com/videos/ -r -p "*.mp4" -o .\videos
.\iod.exe download https://example.com/videos/ -r -p "*.mkv" -o .\videos

# Download all images (JPG, PNG, GIF, etc.)
.\iod.exe download https://example.com/images/ -r -p "*.jpg" -o .\images
.\iod.exe download https://example.com/images/ -r -p "*.png" -o .\images

# Download all audio files
.\iod.exe download https://example.com/music/ -r -p "*.mp3" -o .\music
.\iod.exe download https://example.com/music/ -r -p "*.flac" -o .\music

# Download all text/document files
.\iod.exe download https://example.com/docs/ -r -p "*.txt" -o .\text-files
.\iod.exe download https://example.com/docs/ -r -p "*.doc" -o .\documents
```

### With More Workers (Faster Downloads)

```powershell
# Use 5 concurrent workers for faster downloading
.\iod.exe download https://example.com/pdf/microsoft/ -r -w 5 -o .\downloads

# Use 10 workers (be careful not to overwhelm servers)
.\iod.exe download https://example.com/pdf/microsoft/ -r -w 10 -o .\downloads
```

## Preview Before Downloading

**Always recommended:** List files first to see what you'll be downloading:

```powershell
# List everything recursively
.\iod.exe list https://example.com/pdf/microsoft/ -r

# List with depth limit
.\iod.exe list https://example.com/pdf/microsoft/ -r -d 2

# List with verbose output (shows file sizes)
.\iod.exe list https://example.com/pdf/microsoft/ -r -v

# List only PDFs
.\iod.exe list https://example.com/pdf/microsoft/ -r -p "*.pdf" -v
```

## Directory Structure

The tool **preserves the directory structure** from the website:

```
downloads/
├── brochures/
│   ├── Microsoft_Brochure_198603.pdf
│   └── Microsoft_Product_Line_1983.pdf
├── cpm/
│   ├── Microsoft_BASIC-80_5.0_Reference_1979.pdf
│   └── Microsoft_COBOL-80_1978.pdf
├── mouse/
│   └── Microsoft_Mouse_Programmers_Reference_1989.pdf
└── windows_95/
    └── Inside_Windows_95_1994.pdf
```

## Real-World Examples

### Example 1: Archive All Microsoft Documentation

```powershell
# Download everything, 3 levels deep, 5 workers
.\iod.exe download https://example.com/pdf/microsoft/ -r -d 3 -w 5 -o G:\Archives\Microsoft -v
```

### Example 2: Get Only PDFs from Specific Subdirectory

```powershell
# Download all PDFs from the DOS subdirectories
.\iod.exe download https://example.com/pdf/microsoft/msdos_3.3/ -r -p "*.pdf" -o .\dos-docs
```

### Example 3: Download Everything from Multiple Directories

```powershell
# Download from different sections
.\iod.exe download https://example.com/pdf/microsoft/windows_95/ -r -o .\downloads\windows_95
.\iod.exe download https://example.com/pdf/microsoft/windows_NT_4.0/ -r -o .\downloads\windows_nt
.\iod.exe download https://example.com/pdf/microsoft/visual_studio/ -r -o .\downloads\visual_studio
```

## Important Notes

### ✅ What's Improved

- **✓** Now correctly filters out parent directory links (won't download from `../`)
- **✓** Only downloads files within the specified base URL path
- **✓** Preserves full directory structure
- **✓** Skips already downloaded files (resume support)
- **✓** Handles unlimited depth correctly

### 💡 Best Practices

1. **Preview first**: Always use `list` command to see what will be downloaded
2. **Start small**: Test with `-d 1` or `-d 2` before doing unlimited depth
3. **Use patterns**: Filter by file type to avoid downloading unwanted files
4. **Check space**: Make sure you have enough disk space
5. **Be respectful**: Use reasonable worker counts (3-5) to avoid overwhelming servers
6. **Watch output**: Use `-v` verbose mode to see what's happening

### ⚠️ Warnings

- **Large archives**: Some directories can have hundreds of GB of files
- **Time**: Full recursive downloads can take hours or days for large sites
- **Bandwidth**: Be mindful of your internet connection limits
- **Server load**: Don't use too many workers; respect the server

## Troubleshooting

### Download is slow
```powershell
# Increase workers (but be responsible)
.\iod.exe download URL -r -w 5 -o .\downloads
```

### Too many files
```powershell
# Use depth limit
.\iod.exe download URL -r -d 2 -o .\downloads

# Or filter by type
.\iod.exe download URL -r -p "*.pdf" -o .\downloads
```

### Want to see progress
```powershell
# Use verbose mode
.\iod.exe download URL -r -v -o .\downloads
```

### Resume interrupted download
```powershell
# Just run the same command again - it skips existing files automatically
.\iod.exe download URL -r -o .\downloads
```

## Quick Command Reference

| What You Want | Command |
|---------------|---------|
| Download everything | `.\iod.exe download URL -r -o .\downloads` |
| Download 2 levels deep | `.\iod.exe download URL -r -d 2 -o .\downloads` |
| Download only PDFs | `.\iod.exe download URL -r -p "*.pdf" -o .\downloads` |
| Fast download (5 workers) | `.\iod.exe download URL -r -w 5 -o .\downloads` |
| Preview what's there | `.\iod.exe list URL -r -v` |
| Current dir only (no recursion) | `.\iod.exe download URL -d 0 -o .\downloads` |

---

**The tool is ready for full recursive downloads! Just use the `-r` flag and optionally control the depth with `-d`.** 🚀

