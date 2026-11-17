# 05 - Modern Mathematics, PRNG, and Cryptography

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
Numerical computation and entropy generation require strict separation of concerns. Game mechanics, simulations, and shuffling use Pseudo-Random Number Generators (PRNG), while security tokens, cryptographic salts, and password generation mandate Cryptographically Secure PRNGs (CSPRNG).

---

## 2. PRNG vs CSPRNG in Go
- **`math/rand` (PRNG):** High performance, predictable sequences if seed is known. Never use for security keys. In modern Go, avoid deprecated global `rand.Seed()`; instantiate a dedicated `rand.New(source)` or use `math/rand/v2` (Go 1.22+).
- **`crypto/rand` (CSPRNG):** Reads entropy directly from the OS kernel CSPRNG (`/dev/urandom` on Unix, `CryptGenRandom` / `BCryptGenRandom` on Windows).

---

## 3. High-Precision Monetary Arithmetic (`math/big`)
Standard IEEE-754 floating-point types (`float32`, `float64`) exhibit binary representation drift (e.g. `0.1 + 0.2 != 0.3`). In enterprise finance, use `math/big.Float` or fixed-point integer cents (e.g. `int64` representing micro-cents) to guarantee zero loss of precision.

---

## 4. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
