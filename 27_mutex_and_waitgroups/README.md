# 27 - Mutual Exclusion, RWMutex, and Lock-Free Atomics

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
When multiple Goroutines access the same memory location and at least one access is a write, a **Data Race** occurs unless synchronized.

---

## 2. Synchronization Mechanisms
1. **`sync.Mutex`:**
   Mutual exclusion lock. Only one Goroutine can hold the lock at any time.
2. **`sync.RWMutex`:**
   Reader/Writer mutual exclusion lock. Multiple readers can hold `RLock()` concurrently, but writers require exclusive `Lock()`. Ideal for read-heavy caches and registries.
3. **`sync/atomic` (Lock-Free):**
   Uses atomic CPU hardware instructions (`LOCK XADD`, `CMPXCHG`) to update numeric values and pointer references without mutex sleep or thread yielding overhead.

---

## 3. Detecting Data Races
Go includes an integrated runtime race detector:
```bash
# Execute with race detector active
go test -race ./...
go run -race main.go
```
