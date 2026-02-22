package main

import (
	"fmt"
)

// 1. CONSTANT BLOCKS AND IOTA
// 'iota' is a special Go identifier used to create auto-incrementing enumerations.
// It resets to 0 whenever the 'const' keyword appears in the code.
const (
	Monday    = iota // 0
	Tuesday          // 1 (Go implicitly repeats the '= iota' expression)
	Wednesday        // 2
	Thursday         // 3
	Friday           // 4
	Saturday         // 5
	Sunday           // 6
)

// 2. ADVANCED IOTA: SKIPPING AND BIT SHIFTING
const (
	// The blank identifier '_' throws away the 0 value
	_ = iota 
	
	// iota is now 1. We bit-shift 1 by (10 * 1) bits.
	KB = 1 << (10 * iota) // 1024
	
	// iota is now 2. We bit-shift 1 by (10 * 2) bits.
	MB = 1 << (10 * iota) // 1048576
	
	// iota is now 3. We bit-shift 1 by (10 * 3) bits.
	GB = 1 << (10 * iota) // 1073741824
)

func main() {
	// 3. TYPED VS UNTYPED CONSTANTS
	
	// Untyped Constant: Go just sees a mathematical concept of '42'
	const untypedAge = 42 
	
	// Typed Constant: Go strictly enforces this as an int64
	const typedAge int64 = 42 

	// Because 'untypedAge' is untyped, it can be assigned to ANY numeric type 
	// without needing an explicit conversion.
	var myInt int = untypedAge
	var myFloat float64 = untypedAge
	
	// This would cause a compiler error because typedAge is strictly an int64!
	// var myOtherFloat float64 = typedAge // ERROR: cannot use typedAge (type int64) as type float64

	fmt.Println("--- Typed vs Untyped ---")
	fmt.Printf("myInt: %v (Type: %T)\n", myInt, myInt)
	fmt.Printf("myFloat: %v (Type: %T)\n", myFloat, myFloat)

	// 4. PRINTING IOTA ENUMERATIONS
	fmt.Println("\n--- iota Enumerations ---")
	fmt.Println("Monday is day:", Monday)
	fmt.Println("Wednesday is day:", Wednesday)

	fmt.Println("\n--- iota Bit Shifting ---")
	fmt.Printf("1 KB is %d bytes\n", KB)
	fmt.Printf("1 MB is %d bytes\n", MB)
	fmt.Printf("1 GB is %d bytes\n", GB)
}
