# 06 - Time Engineering & Monotonic Measurements

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
Go approaches time formatting uniquely: instead of abstract pattern tokens like `YYYY-MM-DD HH:mm:ss`, Go uses a specific reference timestamp:
```
Mon Jan 2 15:04:05 MST 2006 (MST is -0700)
```
Memory aid: `1 2 3 4 5 6 -7`
- 1: Month (`01` or `Jan`)
- 2: Day (`02`)
- 3: Hour (`03` for 12hr, `15` for 24hr)
- 4: Minute (`04`)
- 5: Second (`05`)
- 6: Year (`2006` or `06`)
- 7: Timezone offset (`-0700`)

---

## 2. Wall Clock vs Monotonic Clock
- **Wall Clock:** Used for telling current calendar time. Subject to NTP adjustments, daylight savings shifts, and leap seconds.
- **Monotonic Clock:** Continuously increasing clock tick count unaffected by system time adjustments. Go automatically includes monotonic readings in `time.Now()` so operations like `time.Since(start)` or `t2.Sub(t1)` are immune to NTP time jumps.

---

## 3. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
