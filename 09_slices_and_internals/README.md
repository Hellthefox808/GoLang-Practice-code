# 09 - Slices, Capacity Growth, and Memory Internals

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview & Slice Anatomy
While arrays have fixed compile-time lengths, slices are dynamically sized views over underlying backing arrays. In Go, a slice header is a 24-byte structure on 64-bit systems (`SliceHeader`):

```
+-----------------------------------------------+
|                  SliceHeader                  |
+-------------------+-------------+-------------+
| Data Pointer (8B) | Length (8B) | Capacity(8B)|
+-------------------+-------------+-------------+
         |
         v
+--------+--------+--------+--------+--------+
|  elem0 |  elem1 |  elem2 |  elem3 |  elem4 | Backing Array in Memory
+--------+--------+--------+--------+--------+
```

- **`Data`:** Pointer to the first element in the backing array that belongs to the slice.
- **`Len`:** Number of elements currently accessible through the slice (`len(s)`).
- **`Cap`:** Maximum number of elements from the slice's start to the end of the backing array (`cap(s)`).

---

## 2. Capacity Growth Dynamics
When `append()` exceeds the current `cap`:
1. The Go runtime allocates a **new, larger contiguous backing array**.
2. Existing elements are copied to the new memory arena.
3. The slice's `Data` pointer is updated to the new backing array.

### Growth Scaling Algorithm (Modern Go):
- For small slices (< 256 elements), capacity approximately **doubles** ($2\times$).
- For larger slices, capacity grows smoothly at approximately **$1.25\times$ + threshold smoothing**, reducing memory fragmentation.

---

## 3. Sub-Slicing Hazards & Memory Leaks

### Hazard 1: Shared Backing Array Mutation
```go
orig := []string{"A", "B", "C"}
sub := orig[1:3]
sub[0] = "MUTATED"
// orig[1] is now "MUTATED" because both headers point to the same array!
```

### Hazard 2: Memory Retention Leak
If a function returns a tiny 2-byte subslice of a 100MB backing array:
```go
func LeakyFunction() []byte {
    hugePayload := load100MBFile()
    return hugePayload[0:2] // The ENTIRE 100MB array cannot be GC'd!
}
```
**Fix:** Use `slices.Clone()` or `copy()` to detach the slice:
```go
func CleanFunction() []byte {
    hugePayload := load100MBFile()
    return slices.Clone(hugePayload[0:2]) // Detached 2-byte allocation
}
```

---

## 4. Deleting Slice Elements

### Method A: Order-Preserving ($O(N)$)
```go
// Shifts elements after index left by 1 position
slice = append(slice[:index], slice[index+1:]...)
```

### Method B: Unordered Fast Delete ($O(1)$)
```go
// Replace target with the last element, then truncate length
lastIdx := len(slice) - 1
slice[index] = slice[lastIdx]
slice = slice[:lastIdx]
```

---

## 5. Verification & Testing
```bash
go run main.go
go test -v ./...
```
