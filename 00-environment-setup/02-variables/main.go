package main

import (
	"fmt"
)

// 1. PACKAGE-LEVEL VARIABLES
// Variables declared outside of any function can be used anywhere in the package.
// At the package level, you MUST use the 'var' keyword. 
// You can group them together in a 'var' block for cleaner code.
var (
	appVersion string = "1.0.0"
	maxUsers   int    = 100
)

func main() {
	// 2. STANDARD DECLARATION (Explicit Type)
	// Syntax: var [name] [type] = [expression]
	// Use this when you want to be explicit about the type, or when you are 
	// declaring a variable that you will assign a value to later.
	var explicitInt int = 42
	var explicitString string = "Hello, Go!"

	fmt.Println("Explicit Int:", explicitInt)
	fmt.Println("Explicit String:", explicitString)

	// 3. THE "ZERO VALUE"
	// If you declare a variable without assigning a value, Go automatically 
	// assigns it the default "zero value" for its type.
	// - int: 0
	// - float64: 0.0
	// - string: "" (empty string)
	// - bool: false
	// - pointers/slices/maps/channels/functions/interfaces: nil
	var defaultInt int
	var defaultString string
	var defaultBool bool

	fmt.Println("--- Zero Values ---")
	fmt.Println("Default Int:", defaultInt)       // Prints: 0
	fmt.Println("Default String:", defaultString) // Prints: "" (nothing)
	fmt.Println("Default Bool:", defaultBool)     // Prints: false

	// 4. TYPE INFERENCE
	// If you assign an initial value, you can omit the type. 
	// The Go compiler will infer the type based on the value on the right.
	var inferredFloat = 3.14 // Go infers this as a float64
	// The Go compiler will throw an error in case of unused declared variables (extensions, too)
	var unusedInferredBool = true  // Go infers this as a bool

	fmt.Printf("--- Type Inference ---\n")
	// fmt.Printf allows us to format strings. %T prints the data type, %v prints the value.
	fmt.Printf("inferredFloat is of type %T and has value %v\n", inferredFloat, inferredFloat)

	// 5. SHORT VARIABLE DECLARATION (:=)
	// Syntax: [name] := [expression]
	// This is the most common way to declare and initialize variables INSIDE functions.
	// It drops the 'var' keyword entirely. 
	// NOTE: This can ONLY be used inside functions, not at the package level.
	shortMsg := "I was created with the short declaration operator!"
	count := 10

	fmt.Println("--- Short Declarations ---")
	fmt.Println(shortMsg)
	fmt.Println("Count:", count)

	// 6. MULTIPLE ASSIGNMENT
	// Go allows you to declare and assign multiple variables on the same line.
	// This is incredibly useful, especially when functions return multiple values (a core Go feature).
	x, y, z := 1, 2, 3
	fmt.Println("--- Multiple Assignment ---")
	fmt.Println("x:", x, "y:", y, "z:", z)

	// You can also swap variable values without needing a temporary variable
	x, y = y, x 
	fmt.Println("After swap -> x:", x, "y:", y)
}
