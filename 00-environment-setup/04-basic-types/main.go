package main

import (
	"fmt"
	"math/cmplx"
)

// We can declare variables of various types at the package level
var (
	// 1. BOOLEANS
	// Simple true or false. Zero value is false.
	isGolangAwesome bool = true

	// 2. INTEGERS
	// Go has architecture-dependent types (int, uint) and explicitly sized types.
	// 'int' and 'uint' will be 32-bit on 32-bit systems and 64-bit on 64-bit systems.
	// Explicit types: int8, int16, int32, int64 (and their unsigned uint variants).
	maxInt       uint64 = 1<<64 - 1
	standardInt  int    = -42

	// 3. FLOATING POINT & COMPLEX NUMBERS
	// Go has float32 and float64. float64 is the standard and recommended choice.
	pi          float64    = 3.14159265359
	// Go natively supports complex numbers (complex64 and complex128)
	complexNum  complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("--- Booleans and Numbers ---")
	fmt.Printf("Type: %T Value: %v\n", isGolangAwesome, isGolangAwesome)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", complexNum, complexNum)

	// 4. STRINGS, BYTES, AND RUNES
	// A string in Go is a read-only slice of bytes. 
	fmt.Println("\n--- Strings, Bytes, and Runes ---")
	
	// Let's use a string with non-ASCII characters (an emoji and an accented letter)
	myString := "Go! 🚀 caffè"
	fmt.Println("String:", myString)

	// If we get the length of this string using len(), it counts BYTES, not characters!
	// "Go! " = 4 bytes
	// "🚀" = 4 bytes (UTF-8 encoding)
	// "caff" = 4 bytes
	// "è" = 2 bytes (UTF-8 encoding)
	// Total = 14 bytes
	fmt.Printf("Length of string in bytes: %d\n", len(myString))

	// If we loop through the string byte-by-byte, we get raw numbers.
	// 'byte' is actually just an alias for 'uint8'.
	fmt.Print("Bytes: ")
	for i := 0; i < len(myString); i++ {
		fmt.Printf("%x ", myString[i]) // %x prints the hexadecimal byte value
	}
	fmt.Println()

	// To properly count and iterate over actual Unicode characters, we use 'rune'.
	// 'rune' is an alias for 'int32'. It represents a single Unicode code point.
	// We can convert our string to a slice (list) of runes.
	myRunes := []rune(myString)
	
	fmt.Printf("Length of string in runes (actual characters): %d\n", len(myRunes))
	
	fmt.Print("Runes (Unicode Code Points): ")
	// for int range 
	for i := range myRunes {
		// %U prints the standard Unicode format (e.g., U+1F680 for the rocket)
		fmt.Printf("%U ", myRunes[i]) 
	}
	fmt.Println()
}
