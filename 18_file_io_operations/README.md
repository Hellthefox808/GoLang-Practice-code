# 18 - File I/O, Atomic Operations, and Stream Buffering

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
Production file I/O requires avoiding entire-file-in-memory loads for large files, protecting against partial write corruptions, and properly flushing dirty OS caches.

---

## 2. Core Patterns
1. **Modern Standard Library (Go 1.16+):**
   - Deprecated: `io/ioutil.ReadFile`, `io/ioutil.WriteFile`.
   - Modern: `os.ReadFile`, `os.WriteFile`, `io.ReadAll`.
2. **The Atomic Write Pattern:**
   Writing directly to a target file risks corrupting state if the power cuts or process terminates mid-write. The resilient pattern is:
   `CreateTemp` -> Write -> `Sync()` -> `Close()` -> `os.Rename()`. On POSIX and Windows NTFS, file renaming is atomic within the same filesystem volume.
3. **Buffered Streaming with `bufio`:**
   Process gigabyte-scale logs or CSVs in bounded memory footprints using `bufio.Scanner` or `bufio.Reader`.

---

## 3. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
