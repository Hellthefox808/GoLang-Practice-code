# 01 - Hello World & Go Runtime Foundations

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview & Mechanics
In Go, an executable binary starts execution at the `main()` function residing strictly inside `package main`. Unlike C or Java, `main()` accepts no arguments and returns no status codes directly; program termination codes are communicated via `os.Exit(code)` or standard runtime panics.

### The Initialization Sequence:
1. **Package Dependency Graph:** All imported packages are recursively initialized first.
2. **Package-level Variables:** Constant expressions and package-level variables are evaluated.
3. **`init()` Functions:** Any declared `init()` functions execute in lexical order within the compilation unit before `main()`.
4. **`main()` Function:** Execution begins on the primary Goroutine (goroutine 1).

---

## 2. Key Language Concepts Demonstrated
- **Package Contract:** The identifier `package main` instructs the compiler to produce an executable artifact rather than a reusable shared library.
- **IO Decoupling:** Instead of hardcoding `fmt.Println` to the standard operating system descriptor `os.Stdout`, production Go applications accept `io.Writer` interfaces to facilitate unit testing.
- **Runtime Diagnostics:** The `runtime` package provides introspection capabilities to detect CPU cores (`runtime.NumCPU()`), active goroutines (`runtime.NumGoroutine()`), and OS/architecture targets.

---

## 3. How to Run & Verify
```bash
# Direct run via Go toolchain
go run main.go

# Execute unit tests with verbosity
go test -v ./...
```
