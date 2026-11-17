# 28 - Channels, Multiplexing Select, and Worker Pools

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
Go concurrency is rooted in C.A.R. Hoare's **Communicating Sequential Processes (CSP)**:
> *"Do not communicate by sharing memory; instead, share memory by communicating."*

---

## 2. Channel States & Behaviors Matrix
| Operation | Nil Channel | Open Channel | Closed Channel |
|---|---|---|---|
| **Read `<-ch`** | Blocks forever | Returns value (blocks if empty) | Returns zero value, `ok == false` |
| **Write `ch <- v`** | Blocks forever | Sends value (blocks if full) | **PANIC!** |
| **Close `close(ch)`**| **PANIC!** | Closes channel | **PANIC!** |

---

## 3. High-Concurrency Patterns
1. **Directional Types:** Restrict functions to receive-only (`<-chan T`) or send-only (`chan<- T`) to enforce compile-time pipeline safety.
2. **Worker Pool Pattern:** A fixed number of worker Goroutines pull tasks from a shared jobs channel and push to a results channel, bounding memory usage and CPU utilization regardless of spike volume.
3. **`select` Statement:** Multiplexes communication over multiple channels simultaneously with support for timeouts and non-blocking `default:` branches.

---

## 4. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
