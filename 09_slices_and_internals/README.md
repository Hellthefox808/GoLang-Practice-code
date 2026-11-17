# 09 - Slices, Capacity Growth, and Memory Internals

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
A slice is a dynamically sized view into the elements of an underlying array. Under the hood, a slice is a 24-byte (on 64-bit architecture) struct header defined conceptually as:
```go
type SliceHeader struct {
    Data uintptr // pointer to backing array
    Len  int     // current count of elements
    Cap  int     // maximum capacity before reallocation
}
```

---

## 2. Growth Dynamics
When `append()` exceeds `cap`, the Go runtime allocates a new, larger backing array, copies the existing elements into it, and updates the pointer.
- For small slices (< 256 elements in modern Go), capacity typically doubles.
- For larger slices, growth scales smoothly around `1.25x` to balance allocation speed with memory fragmentation.

---

## 3. Sub-Slicing Hazards
Creating a sub-slice (`s[a:b]`) creates a new slice header that still points to the **same backing array**. Modifying elements in the sub-slice directly modifies the original slice unless cloned via `slices.Clone()` or copied via `copy()`.

---

## 4. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
