# How Recursive Downloads Work

## Visual Guide

### Directory Structure on Server

```
https://example.com/pdf/microsoft/
│
├── file1.pdf                    ← Downloaded with -d 0
├── file2.txt                    ← Downloaded with -d 0
│
├── brochures/                   ← Entered with -r -d 1 or higher
│   ├── brochure1.pdf            ← Downloaded with -d 1
│   └── brochure2.pdf            ← Downloaded with -d 1
│
├── windows_95/                  ← Entered with -r -d 1 or higher
│   ├── book1.pdf                ← Downloaded with -d 1
│   ├── book2.pdf                ← Downloaded with -d 1
│   │
│   └── drivers/                 ← Entered with -r -d 2 or higher
│       ├── driver1.exe          ← Downloaded with -d 2
│       └── driver2.exe          ← Downloaded with -d 2
│       │
│       └── extras/              ← Entered with -r -d 3 or higher
│           ├── extra1.zip       ← Downloaded with -d 3
│           └── extra2.zip       ← Downloaded with -d 3
│
└── msdos/                       ← Entered with -r -d 1 or higher
    ├── version_3.3/             ← Entered with -r -d 2 or higher
    │   └── manual.pdf           ← Downloaded with -d 2
    │
    └── version_5.0/             ← Entered with -r -d 2 or higher
        └── guide.pdf            ← Downloaded with -d 2
```

## Command Comparison

### No Recursion (`-d 0` or no `-r`)
```powershell
.\iod.exe download https://example.com/pdf/microsoft/ -d 0 -o .\downloads
```
**Downloads:**
- ✓ file1.pdf
- ✓ file2.txt
- ✗ Nothing from subdirectories

**Result:** 2 files

---

### Depth 1 (`-r -d 1`)
```powershell
.\iod.exe download https://example.com/pdf/microsoft/ -r -d 1 -o .\downloads
```
**Downloads:**
- ✓ file1.pdf
- ✓ file2.txt
- ✓ brochures/brochure1.pdf
- ✓ brochures/brochure2.pdf
- ✓ windows_95/book1.pdf
- ✓ windows_95/book2.pdf
- ✓ msdos/ (immediate files only)
- ✗ windows_95/drivers/ (too deep)
- ✗ msdos/version_3.3/ (too deep)

**Result:** ~180+ files (all files in immediate subdirectories)

---

### Depth 2 (`-r -d 2`)
```powershell
.\iod.exe download https://example.com/pdf/microsoft/ -r -d 2 -o .\downloads
```
**Downloads:**
- ✓ Everything from depth 1
- ✓ windows_95/drivers/driver1.exe
- ✓ windows_95/drivers/driver2.exe
- ✓ msdos/version_3.3/manual.pdf
- ✓ msdos/version_5.0/guide.pdf
- ✗ windows_95/drivers/extras/ (too deep)

**Result:** More files (2 levels deep)

---

### Unlimited Depth (`-r` without `-d`, or `-r -d -1`)
```powershell
.\iod.exe download https://example.com/pdf/microsoft/ -r -o .\downloads
```
**Downloads:**
- ✓ **EVERYTHING** from all subdirectories
- ✓ windows_95/drivers/extras/extra1.zip
- ✓ All files at any depth

**Result:** Complete mirror of the entire directory tree

---

## How the Tool Traverses Directories

```
1. Start at: https://example.com/pdf/microsoft/
   │
   ├─→ Parse HTML page
   │
   ├─→ Find all links
   │   ├─→ Links ending with / = Directories
   │   └─→ Links without / = Files
   │
   ├─→ Download all files at current level
   │
   └─→ If -r flag is set:
       └─→ For each directory found:
           ├─→ Check if depth limit allows
           ├─→ Enter directory (recursive call)
           └─→ Repeat process
```

## Local Directory Structure After Download

```
downloads/
│
├── file1.pdf                    ← Preserves original structure
├── file2.txt
│
├── brochures/
│   ├── brochure1.pdf
│   └── brochure2.pdf
│
├── windows_95/
│   ├── book1.pdf
│   ├── book2.pdf
│   └── drivers/
│       ├── driver1.exe
│       ├── driver2.exe
│       └── extras/
│           ├── extra1.zip
│           └── extra2.zip
│
└── msdos/
    ├── version_3.3/
    │   └── manual.pdf
    └── version_5.0/
        └── guide.pdf
```

## Examples with Expected Results

### Example 1: Just One Subdirectory Deep
```powershell
.\iod.exe download https://example.com/pdf/microsoft/windows_95/ -r -d 1 -o .\win95
```
**Gets:**
- All files in windows_95/
- All files in immediate subdirectories of windows_95/
- Does NOT go deeper than 1 level

### Example 2: Everything Under windows_95/
```powershell
.\iod.exe download https://example.com/pdf/microsoft/windows_95/ -r -o .\win95
```
**Gets:**
- Every single file under windows_95/ at any depth

### Example 3: Only PDFs, All Depths
```powershell
.\iod.exe download https://example.com/pdf/microsoft/ -r -p "*.pdf" -o .\pdfs
```
**Gets:**
- All .pdf files from microsoft/ and all subdirectories
- Skips .txt, .zip, .exe, etc.

## Tips for Understanding Depth

Think of depth as "how many levels down from the starting URL":

- **Depth 0**: Starting URL only (no subdirectories)
- **Depth 1**: Starting URL + 1 folder down
- **Depth 2**: Starting URL + 2 folders down
- **Depth 3**: Starting URL + 3 folders down
- **Depth -1** (or unlimited): All the way down, no matter how deep

## Performance Considerations

| Depth | Approximate Files* | Download Time* |
|-------|-------------------|----------------|
| 0 | 5-10 | Seconds |
| 1 | 100-200 | Minutes |
| 2 | 500-1000 | 30-60 min |
| 3+ | 1000-5000+ | Hours |
| Unlimited | 5000-50000+ | Hours to days |

*Estimates for example.com/pdf/microsoft/

## Smart Download Strategy

1. **Preview first:**
   ```powershell
   .\iod.exe list URL -r -d 1
   ```

2. **Start shallow:**
   ```powershell
   .\iod.exe download URL -r -d 1 -o .\downloads
   ```

3. **Increase depth gradually:**
   ```powershell
   .\iod.exe download URL -r -d 2 -o .\downloads
   ```

4. **Go unlimited when ready:**
   ```powershell
   .\iod.exe download URL -r -o .\downloads
   ```

---

**Now you understand exactly how recursive downloads work! 🎓**

