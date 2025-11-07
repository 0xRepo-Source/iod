# GitHub Repository Setup Guide

## Quick Setup

### 1. Create GitHub Repository

1. Go to https://github.com/new
2. Repository name: `iod` (or `indexof-downloader`)
3. Description: "Universal file downloader for directory listings - Download any file type from Apache/Nginx index pages"
4. Choose: Public
5. **DO NOT** initialize with README, .gitignore, or license (we already have them)
6. Click "Create repository"

### 2. Connect Local Repository to GitHub

```bash
# Set your repository URL (replace YOUR-USERNAME)
git remote add origin https://github.com/YOUR-USERNAME/iod.git

# Verify remote
git remote -v

# Set default branch name
git branch -M main

# Push to GitHub
git push -u origin main
```

### 3. First Commit

```bash
# Commit all files
git commit -m "Initial release v1.0.0 - Universal file downloader for directory listings"

# Push to GitHub
git push -u origin main
```

### 4. Create First Release (Optional)

```bash
# Create and push a tag
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

Then on GitHub:
1. Go to your repository
2. Click "Releases" → "Create a new release"
3. Choose tag: v1.0.0
4. Release title: "v1.0.0 - Initial Release"
5. Description: Copy from CHANGELOG.md
6. Attach binaries (or use GitHub Actions to auto-build)
7. Click "Publish release"

## Repository Settings

### Topics (Tags)
Add these topics to help people find your project:
- `go`
- `golang`
- `downloader`
- `cli`
- `directory-listing`
- `apache-index`
- `file-downloader`
- `recursive-download`
- `bitsavers`
- `archive`

### About Section
**Description:** Universal file downloader for directory listings - Download any file type from Apache/Nginx index pages

**Website:** (optional - your docs site if you have one)

**Topics:** Add the topics listed above

### Features to Enable
- ✅ Issues
- ✅ Projects (optional)
- ✅ Discussions (optional)
- ✅ Actions (for automated builds)

## GitHub Actions Setup

The repository includes two workflow files:

1. **`.github/workflows/build.yml`** - Builds binaries for all platforms on tag push
2. **`.github/workflows/test.yml`** - Runs tests on every push/PR

These will activate automatically when you push to GitHub.

## Complete Git Setup Commands

```bash
# Initialize (already done)
git init

# Stage all files (already done)
git add .

# Commit
git commit -m "Initial release v1.0.0 - Universal file downloader for directory listings

Features:
- Universal file downloader for directory listings
- Recursive directory traversal with depth control
- Pattern matching for any file type
- Concurrent downloads with progress bars
- Works with Apache/Nginx index pages
- Comprehensive documentation"

# Add remote (replace YOUR-USERNAME)
git remote add origin https://github.com/YOUR-USERNAME/iod.git

# Push to GitHub
git branch -M main
git push -u origin main

# Create release tag
git tag -a v1.0.0 -m "Release v1.0.0 - Initial Release"
git push origin v1.0.0
```

## README Badges (Optional)

Add these to the top of README.md after pushing to GitHub:

```markdown
![Go Version](https://img.shields.io/github/go-mod/go-version/YOUR-USERNAME/iod)
![License](https://img.shields.io/github/license/YOUR-USERNAME/iod)
![Release](https://img.shields.io/github/v/release/YOUR-USERNAME/iod)
![Downloads](https://img.shields.io/github/downloads/YOUR-USERNAME/iod/total)
```

## What's Already Prepared

✅ `.gitignore` - Properly configured  
✅ `LICENSE` - MIT License  
✅ `README.md` - Comprehensive documentation  
✅ `CONTRIBUTING.md` - Contribution guidelines  
✅ `CHANGELOG.md` - Version history  
✅ `.github/workflows/` - CI/CD pipelines  
✅ All documentation files  
✅ Go module configuration  

## After Pushing to GitHub

1. **Enable GitHub Actions** (should be automatic)
2. **Add topics** to your repository
3. **Create a release** for v1.0.0
4. **Star your own repo** to get it started 😊
5. **Share on social media** or relevant forums

## Repository Description Template

```
🚀 IOD - Universal file downloader for directory listings

Download ANY file type from Apache/Nginx directory index pages with recursive traversal, 
pattern matching, concurrent downloads, and progress tracking. Perfect for archiving 
old software repositories, downloading ISO collections, or mirroring documentation sites.

Features: Recursive downloads, depth control, pattern filtering, concurrent workers, 
progress bars, resume support. Works with PDFs, ZIPs, ISOs, videos, and any file type.
```

---

**Ready to push to GitHub!** 🎉

Just replace `YOUR-USERNAME` with your GitHub username in the commands above.
