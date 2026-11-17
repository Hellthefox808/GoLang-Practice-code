# 26 - Goroutines, Scheduler Dynamics, and Bounded Concurrency

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
A Goroutine is an extremely lightweight thread of execution managed by the Go runtime, not the operating system kernel.

---

## 2. The M:N Scheduler Architecture
- **G (Goroutine):** Contains stack pointer, program counter, and scheduling metadata. Starts with only 2KB of memory.
- **M (Machine):** An OS kernel thread created and managed by the operating system.
- **P (Processor):** A logical processor representing a context required to execute Go code (defaults to `runtime.GOMAXPROCS()`, equal to CPU core count).

---

## 3. Concurrency Hygiene Rules
1. **`wg.Add(1)` strictly before `go func()`:** Calling `Add(1)` inside the spawned goroutine introduces an immediate race condition where `wg.Wait()` may return before the goroutine even increments the counter.
2. **Prevent Socket/FD Leaks:** When performing concurrent HTTP requests, always read the body (`io.Copy(io.Discard, resp.Body)`) and close it. Failing to drain the body leaks OS socket file descriptors.
3. **Bounded Semaphore Pattern:** Launching millions of unbounded goroutines to hit external systems can exhaust socket limits and trigger rate limiting. Use buffered channels (`make(chan struct{}, N)`) to enforce maximum concurrency bounds.

---

## 4. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
