# Universal File Type Support

## 🌟 Works with ALL File Types!

IndexOf Downloader is a **universal downloader** that works with **any file extension**. It doesn't care what type of file you're downloading - if it's on a directory listing page, it can download it!

## Supported File Types (All of them!)

### 📄 Documents
```powershell
# PDFs
.\iod.exe download URL -r -p "*.pdf" -o .\pdfs

# Word Documents
.\iod.exe download URL -r -p "*.doc" -o .\documents
.\iod.exe download URL -r -p "*.docx" -o .\documents

# Excel Spreadsheets
.\iod.exe download URL -r -p "*.xls" -o .\spreadsheets
.\iod.exe download URL -r -p "*.xlsx" -o .\spreadsheets

# PowerPoint Presentations
.\iod.exe download URL -r -p "*.ppt" -o .\presentations
.\iod.exe download URL -r -p "*.pptx" -o .\presentations

# Text files
.\iod.exe download URL -r -p "*.txt" -o .\text
.\iod.exe download URL -r -p "*.md" -o .\markdown
.\iod.exe download URL -r -p "*.rtf" -o .\documents
```

### 📦 Archives & Compressed Files
```powershell
# ZIP files
.\iod.exe download URL -r -p "*.zip" -o .\archives

# RAR files
.\iod.exe download URL -r -p "*.rar" -o .\archives

# 7-Zip files
.\iod.exe download URL -r -p "*.7z" -o .\archives

# TAR archives
.\iod.exe download URL -r -p "*.tar" -o .\archives
.\iod.exe download URL -r -p "*.tar.gz" -o .\archives
.\iod.exe download URL -r -p "*.tgz" -o .\archives

# Other compressed formats
.\iod.exe download URL -r -p "*.gz" -o .\archives
.\iod.exe download URL -r -p "*.bz2" -o .\archives
```

### 💿 Disk Images
```powershell
# ISO images
.\iod.exe download URL -r -p "*.iso" -o .\isos

# IMG files
.\iod.exe download URL -r -p "*.img" -o .\images

# VHD/VMDK (Virtual disks)
.\iod.exe download URL -r -p "*.vhd" -o .\virtual-disks
.\iod.exe download URL -r -p "*.vmdk" -o .\virtual-disks
```

### 💻 Software & Executables
```powershell
# Windows executables
.\iod.exe download URL -r -p "*.exe" -o .\software

# MSI installers
.\iod.exe download URL -r -p "*.msi" -o .\installers

# DLL files
.\iod.exe download URL -r -p "*.dll" -o .\libraries

# Linux binaries
.\iod.exe download URL -r -p "*.deb" -o .\packages
.\iod.exe download URL -r -p "*.rpm" -o .\packages

# AppImage
.\iod.exe download URL -r -p "*.AppImage" -o .\software
```

### 🎬 Video Files
```powershell
# MP4 videos
.\iod.exe download URL -r -p "*.mp4" -o .\videos

# AVI videos
.\iod.exe download URL -r -p "*.avi" -o .\videos

# MKV videos
.\iod.exe download URL -r -p "*.mkv" -o .\videos

# MOV videos
.\iod.exe download URL -r -p "*.mov" -o .\videos

# WebM videos
.\iod.exe download URL -r -p "*.webm" -o .\videos

# FLV videos
.\iod.exe download URL -r -p "*.flv" -o .\videos

# MPEG/MPG videos
.\iod.exe download URL -r -p "*.mpeg" -o .\videos
.\iod.exe download URL -r -p "*.mpg" -o .\videos
```

### 🎵 Audio Files
```powershell
# MP3 audio
.\iod.exe download URL -r -p "*.mp3" -o .\music

# FLAC audio
.\iod.exe download URL -r -p "*.flac" -o .\music

# WAV audio
.\iod.exe download URL -r -p "*.wav" -o .\audio

# OGG/Vorbis
.\iod.exe download URL -r -p "*.ogg" -o .\audio

# AAC/M4A
.\iod.exe download URL -r -p "*.aac" -o .\audio
.\iod.exe download URL -r -p "*.m4a" -o .\audio

# WMA
.\iod.exe download URL -r -p "*.wma" -o .\audio
```

### 🖼️ Images
```powershell
# JPEG images
.\iod.exe download URL -r -p "*.jpg" -o .\images
.\iod.exe download URL -r -p "*.jpeg" -o .\images

# PNG images
.\iod.exe download URL -r -p "*.png" -o .\images

# GIF images
.\iod.exe download URL -r -p "*.gif" -o .\images

# BMP images
.\iod.exe download URL -r -p "*.bmp" -o .\images

# TIFF images
.\iod.exe download URL -r -p "*.tiff" -o .\images
.\iod.exe download URL -r -p "*.tif" -o .\images

# WebP images
.\iod.exe download URL -r -p "*.webp" -o .\images

# SVG images
.\iod.exe download URL -r -p "*.svg" -o .\images

# Raw image formats
.\iod.exe download URL -r -p "*.raw" -o .\images
.\iod.exe download URL -r -p "*.cr2" -o .\images
```

### 💾 Data & Database Files
```powershell
# SQL dumps
.\iod.exe download URL -r -p "*.sql" -o .\databases

# Database files
.\iod.exe download URL -r -p "*.db" -o .\databases
.\iod.exe download URL -r -p "*.sqlite" -o .\databases
.\iod.exe download URL -r -p "*.mdb" -o .\databases

# CSV files
.\iod.exe download URL -r -p "*.csv" -o .\data

# JSON files
.\iod.exe download URL -r -p "*.json" -o .\data

# XML files
.\iod.exe download URL -r -p "*.xml" -o .\data
```

### 👨‍💻 Source Code & Development
```powershell
# Source code archives
.\iod.exe download URL -r -p "*.tar.gz" -o .\source-code

# Python files
.\iod.exe download URL -r -p "*.py" -o .\python

# JavaScript files
.\iod.exe download URL -r -p "*.js" -o .\javascript

# Java files
.\iod.exe download URL -r -p "*.java" -o .\java
.\iod.exe download URL -r -p "*.jar" -o .\jars

# C/C++ files
.\iod.exe download URL -r -p "*.c" -o .\c-code
.\iod.exe download URL -r -p "*.cpp" -o .\cpp-code
.\iod.exe download URL -r -p "*.h" -o .\headers

# Go files
.\iod.exe download URL -r -p "*.go" -o .\go-code
```

### 🎮 Game Files & ROMs
```powershell
# ROM files
.\iod.exe download URL -r -p "*.rom" -o .\roms
.\iod.exe download URL -r -p "*.nes" -o .\roms
.\iod.exe download URL -r -p "*.gba" -o .\roms
.\iod.exe download URL -r -p "*.gb" -o .\roms

# Game packages
.\iod.exe download URL -r -p "*.pak" -o .\game-files
.\iod.exe download URL -r -p "*.wad" -o .\game-files
```

### 📧 Email & Communication
```powershell
# Email files
.\iod.exe download URL -r -p "*.eml" -o .\emails
.\iod.exe download URL -r -p "*.msg" -o .\emails

# PST files
.\iod.exe download URL -r -p "*.pst" -o .\mailboxes
```

### 🔐 Certificates & Keys
```powershell
# Certificates
.\iod.exe download URL -r -p "*.pem" -o .\certificates
.\iod.exe download URL -r -p "*.crt" -o .\certificates
.\iod.exe download URL -r -p "*.cer" -o .\certificates

# Keys
.\iod.exe download URL -r -p "*.key" -o .\keys
```

### 🌐 Web Files
```powershell
# HTML files
.\iod.exe download URL -r -p "*.html" -o .\web-pages

# CSS files
.\iod.exe download URL -r -p "*.css" -o .\stylesheets

# Web archives
.\iod.exe download URL -r -p "*.warc" -o .\web-archives
```

## 🎯 Real-World Use Cases

### Download All Software from Archive
```powershell
# Get all EXE and MSI installers
.\iod.exe download https://example.com/software/ -r -p "*.exe" -o .\software
.\iod.exe download https://example.com/software/ -r -p "*.msi" -o .\installers
```

### Archive Old FTP Site
```powershell
# Download everything (all file types)
.\iod.exe download http://old-ftp-site.com/pub/ -r -o .\ftp-archive
```

### Download ISO Collection
```powershell
# Get all operating system ISOs
.\iod.exe download https://mirror.example.com/isos/ -r -p "*.iso" -o .\isos
```

### Download Music Library
```powershell
# Download all MP3 and FLAC files
.\iod.exe download https://example.com/music/ -r -p "*.mp3" -o .\music
.\iod.exe download https://example.com/music/ -r -p "*.flac" -o .\music-flac
```

### Download Research Data
```powershell
# Download all CSV data files
.\iod.exe download https://data.example.com/datasets/ -r -p "*.csv" -o .\datasets
```

## 💡 Tips

1. **Any extension works**: The tool doesn't validate or restrict file types
2. **Case-insensitive**: `*.PDF` works the same as `*.pdf`
3. **Wildcards**: Use `*` to match any characters
4. **No pattern = download all**: Omit `-p` to download every file type
5. **Combine commands**: Run multiple commands with different patterns to organize by type

## 🚀 The Bottom Line

**IndexOf Downloader downloads ANY file type!** It's a universal downloader for directory listings. Whether you're downloading:
- Documents (PDF, DOC, XLS)
- Archives (ZIP, RAR, 7Z, TAR)
- Images (JPG, PNG, GIF)
- Videos (MP4, AVI, MKV)
- Audio (MP3, FLAC, WAV)
- Software (EXE, MSI, DEB, RPM)
- ISOs, ROMs, or anything else...

**It just works!** 🎉

