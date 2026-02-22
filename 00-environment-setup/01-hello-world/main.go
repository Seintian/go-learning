// 1. Package Declaration
// Every Go source file must start with a package declaration.
// Packages are Go's way of organizing code.
// The name "main" is special: it tells the Go compiler that this file should be
// compiled as a standalone executable program rather than a shared library.
package main

// 2. Import Statements
// The import keyword is used to include code from other packages.
// "fmt" stands for "format" and is part of Go's Standard Library.
// It contains functions for formatting text, reading input, and printing to the console.
import (
	"fmt"
)

// 3. The Main Function
// The main() function is the entry point of the compiled program.
// When you execute the program, the code inside this function runs first.
// Notice that the opening brace `{` must be on the same line as the function declaration.
func main() {
	// Collect both return values from fmt.Println
	var n int
	var err error

	// fmt.Println outputs the string to the standard output (your terminal)
	// and automatically appends a newline character at the end.
	n, err = fmt.Println("Hello, World! My Go environment is successfully set up.")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("What were n and err? n:", n, "err:", err)

	// NOTE: There's no ternary expression in Go, so here's the boring equivalent if-else
	var message string
	if n == 56 {
		message = "They were the bytes written (+ nul) to the standard output and the eventual error."
	} else {
		message = "They were something else."
	}
	fmt.Println(message)
}
