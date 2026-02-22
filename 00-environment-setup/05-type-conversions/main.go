package main

import (
	"fmt"
	"strconv" // Crucial standard library package for string conversions
)

func main() {
	// 1. NUMERIC CONVERSIONS
	// The general syntax for conversion is T(v) where T is the Type and v is the Value.
	var x int = 10
	var y float64 = 3.14

	// This would cause a compiler error: "invalid operation: x + y (mismatched types int and float64)"
	// var result = x + y 

	// We must explicitly convert x to a float64 first
	var result float64 = float64(x) + y
	fmt.Println("--- Numeric Conversions ---")
	fmt.Printf("Result: %v (Type: %T)\n", result, result)

	// When converting floating-point numbers to integers, the decimal is TRUNCATED (not rounded).
	var roundedDown int = int(y) // 3.14 becomes 3
	fmt.Printf("Truncated Float: %v (Type: %T)\n", roundedDown, roundedDown)

	// 2. STRINGS TO SLICES (Bytes and Runes)
	// As we saw in the last lesson, strings are fundamentally slices of bytes.
	// We can cast freely between strings and slices of bytes or runes.
	fmt.Println("\n--- Strings and Slices ---")
	myStr := "Hello"

	// Convert string to a slice of bytes
	byteSlice := []byte(myStr)
	fmt.Printf("String to []byte: %x (%s)\n", byteSlice, byteSlice) // Prints the ASCII/UTF-8 byte values

	// Convert slice of bytes back to a string
	restoredStr := string(byteSlice)
	fmt.Printf("[]byte back to String: %s\n", restoredStr)

	// 3. THE STRCONV PACKAGE (String Conversions)
	// A massive trap for beginners is trying to convert an integer to a string like this:
	// wrongStr := string(65) 
	// This does NOT produce "65". It produces "A" (because 65 is the Unicode value for 'A').
	// To convert numbers to text representations, you MUST use the 'strconv' package.

	fmt.Println("\n--- The strconv Package ---")

	// Itoa: Integer to ASCII (String)
	age := 42
	ageStr := strconv.Itoa(age)
	fmt.Printf("Itoa Conversion: %q (Type: %T)\n", ageStr, ageStr) // %q prints with quotes

	// Atoi: ASCII (String) to Integer
	// This function returns TWO values: the result, and an error if the string wasn't a valid number.
	// We will cover error handling deeply later, so we use '_' to ignore the error for now.
	strPrice := "250"
	priceInt, _ := strconv.Atoi(strPrice)
	fmt.Printf("Atoi Conversion: %v (Type: %T)\n", priceInt, priceInt)

	// Parsing floats from strings
	strPi := "3.14159"
	// The '64' means we want to parse it into a float64
	parsedFloat, _ := strconv.ParseFloat(strPi, 64) 
	fmt.Printf("ParseFloat Conversion: %v (Type: %T)\n", parsedFloat, parsedFloat)
}
