# Chapter 00: Environment Setup & Fundamentals

Welcome to your first step in mastering Go! Before diving into the syntax, it is crucial to establish a solid local development environment and understand how Go manages projects and executes code. 

## 1. Installing Go

Go provides pre-compiled binaries for all major operating systems.

1. **Download:** Navigate to the official [Go Downloads page](https://go.dev/dl/) and download the installer for your operating system (for Windows, this will be an `.msi` file).
2. **Install:** Run the installer and follow the prompts. The default installation path (`C:\Program Files\Go` on Windows) is recommended. The installer will automatically add the Go binary directory to your system's `PATH` environment variable.
3. **Verify:** Open a new terminal (Command Prompt or PowerShell) and run:
   `go version`
   If the installation was successful, the terminal will print the installed version of Go and your operating system/architecture.

## 2. Setting Up Your Editor

While you can write Go in any text editor, using an IDE with strong Go support will drastically improve your learning experience by providing code formatting, linting, and autocomplete.

**Recommended Setup: Visual Studio Code (VS Code)**
1. Download and install [VS Code](https://code.visualstudio.com/).
2. Open VS Code and navigate to the Extensions view.
3. Search for **"Go"** and install the official extension provided by the **Go Team at Google**.
4. Once installed, open any `.go` file or open the Command Palette, type `Go: Install/Update Tools`, select all the available tools, and click OK. This installs the language server and debugging tools.

## 3. Understanding Go Modules (Modern Dependency Management)

Historically, Go required all code to be written inside a specific directory called the `GOPATH`. Modern Go (version 1.11+) uses **Go Modules**, which allows you to place your project anywhere on your system.

To initialize a new Go module in your main repository folder, open your terminal at the root of `GO-LEARNING` and run:

`go mod init github.com/yourusername/go-learning`

*(Note: The module path is typically a URL representing where the code is hosted, but it can be any unique string if you are just learning locally).*

This creates a `go.mod` file, which tracks your project's Go version and any external dependencies.

## 4. How Go Code is Executed

Go is a compiled language, meaning source code must be translated into machine code before it runs. However, Go's tooling makes this process incredibly fast.

Navigate into any of the subdirectories in this chapter and you can use two primary commands:

* **`go run main.go`**: This command compiles the code into a temporary directory, executes it, and then deletes the executable. It is perfect for rapid testing and learning.
* **`go build main.go`**: This command compiles the code and generates a permanent, standalone executable file in your current directory. You can then run this file directly.

## 5. Chapter Overview

This directory contains several nested projects. Navigate into each folder to explore the `.go` files and run them. 

* **`01-hello-world`**: The structure of a Go program, the `main` package, imports, and printing to the console.
* **`02-variables`**: Declaring variables, the `var` keyword, short variable declarations (`:=`), and Go's concept of "zero values".
* **`03-constants`**: Declaring immutable values, untyped vs. typed constants, and utilizing the `iota` identifier for enumerations.
* **`04-basic-types`**: Exploring Go's built-in primitive types (Strings, Integers, Floats, Booleans, Bytes, and Runes).
* **`05-type-conversions`**: How to safely and explicitly convert between different data types.
* **`06-operators`**: Using arithmetic, comparison, logical, and bitwise operators.
