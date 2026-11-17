# 10 - Hash Tables, Map Semantics, and Safety

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview & Hash Table Architecture
Go maps are implemented via a dynamically resizable hash table (`runtime.hmap`). Memory is arranged in an array of **Buckets (`bmap`)**, where each bucket stores up to **8 key-value pairs**:

```
+-------------------------------------------------------+
|                       hmap Header                     |
|  count: int | B: uint8 (2^B buckets) | buckets: *bmap |
+-------------------------------------------------------+
                           |
                           v
        Bucket 0                     Bucket 1
+-----------------------+   +-----------------------+
| tophash: [8]uint8     |   | tophash: [8]uint8     |
| keys   : [8]KeyType   |   | keys   : [8]KeyType   |
| values : [8]ValType   |   | values : [8]ValType   |
| overflow: *bmap       |   | overflow: *bmap       |
+-----------------------+   +-----------------------+
```

1. **Hashing:** Go hashes the key using an architecture-optimized hash algorithm (AES-NI hardware accelerated on x86_64).
2. **Bucket Selection:** The lowest `B` bits determine the target bucket.
3. **Key Comparison:** The top 8 bits (`tophash`) provide fast byte-level filtering before executing full equality checks.

---

## 2. Fundamental Safety Rules

### Rule 1: The Nil Map Trap
- Reading from a `nil` map is completely safe and evaluates to the zero-value:
  ```go
  var m map[string]int
  fmt.Println(m["missing"]) // Prints 0 (Safe)
  ```
- Writing to a `nil` map causes an **immediate, unrecoverable runtime panic**:
  ```go
  var m map[string]int
  m["key"] = 100 // PANIC: assignment to entry in nil map
  ```
  *Remedy:* Always initialize maps with `make(map[K]V, hint)` or `{}`.

### Rule 2: The Comma-Ok Idiom
Because reading a missing key yields the type's zero-value, `m[key]` cannot distinguish between an absent key and a key whose stored value is `0` or `""`:
```go
val, ok := registry[key]
if !ok {
    // Key does not exist
}
```

### Rule 3: Iteration Non-Determinism
Go intentionally randomizes the starting bucket offset when executing `for k, v := range m`. This prevents applications from depending on consistent hash table iteration order and protects against HashDOS attacks.

### Rule 4: Concurrency Hazard
Go maps are **NOT thread-safe**. Concurrent reads and writes without synchronization trigger an immediate unrecoverable crash:
```
fatal error: concurrent map writes
```
Use `sync.RWMutex` (see [Module 27](../27_mutex_and_waitgroups)) or `sync.Map` for concurrent workloads.

---

## 3. Verification & Testing
```bash
go run main.go
go test -v ./...
```
