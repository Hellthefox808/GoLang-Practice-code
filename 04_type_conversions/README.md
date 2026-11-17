# 04 - Type Conversions, Parsing, and Safe Casts

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
In Go, there is **no implicit type conversion** (coercion). Even assigning an `int32` to an `int64` requires an explicit type cast: `int64(myInt32)`. This strict rule eliminates hidden truncation, sign-extension, and floating-point rounding bugs common in other languages.

---

## 2. Core Conversion Mechanics
1. **Numeric Casting:**
   - Casting wider types to narrower types (e.g., `int64` to `int32`) silently truncates high bits unless bounds-checked manually.
   - Float-to-integer conversion discards decimal values towards zero (truncation).
2. **`strconv` Package for Text Representations:**
   - Parsing: `strconv.Atoi`, `strconv.ParseInt`, `strconv.ParseFloat`, `strconv.ParseBool`.
   - Formatting: `strconv.Itoa`, `strconv.FormatInt`, `strconv.FormatFloat`, `strconv.FormatBool`.
3. **Strings and Byte Slices:**
   - Converting `string` to `[]byte` allocates a new backing slice in memory because Go strings are immutable, while slices are mutable.

---

## 3. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
