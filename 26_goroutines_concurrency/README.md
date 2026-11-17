# 26 - Goroutines, Scheduler Dynamics, and Bounded Concurrency

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview & Scheduler Mechanics
Goroutines are cooperative, lightweight user-space threads managed entirely by the Go runtime scheduler.

### Goroutine vs. Operating System Thread:
| Attribute | Goroutine (`G`) | OS Thread (`M`) |
|---|---|---|
| **Initial Stack Size** | **2 KB** (dynamic grow/shrink) | **1 MB – 8 MB** (fixed) |
| **Creation Overhead** | ~few nanoseconds | ~1–2 microseconds |
| **Context Switch Cost** | User-space (~10–20 ns) | Kernel transition (~1000–1500 ns) |
| **Scalability** | Millions per process | Thousands per process |

---

## 2. The Golden Rules of `sync.WaitGroup`

```
CORRECT:                                WRONG (RACE CONDITION):
for _, u := range urls {                for _, u := range urls {
    wg.Add(1)  // <--- BEFORE 'go'          go func() {
    go func() {                                 wg.Add(1) // <--- CRASH/RACE!
        defer wg.Done()                             defer wg.Done()
    }()                                     }()
}                                       }
wg.Wait()                               wg.Wait()
```

If `wg.Add(1)` is placed inside the goroutine, `wg.Wait()` can execute on the main thread and terminate the process before the runtime scheduler even allocates CPU time to the new goroutines!

---

## 3. Bounded Concurrency: The Counting Semaphore Pattern
Launching unbounded goroutines against network endpoints exhausts operating system socket descriptors and triggers rate-limiting bans.

A buffered channel of size `N` functions as a hardware-efficient **counting semaphore**:
```go
semaphore := make(chan struct{}, maxParallel) // e.g. 3

for _, url := range urls {
    wg.Add(1)
    go func(target string) {
        semaphore <- struct{}{}        // Acquire token (blocks if 3 workers are active)
        defer func() {
            <-semaphore                // Release token
            wg.Done()
        }()

        probe(target)
    }(url)
}
wg.Wait()
```

---

## 4. Verification & Testing
```bash
go run main.go
go test -v ./...
# Verify data race freedom
go test -race ./...
```
