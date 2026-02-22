package main

import (
	"fmt"
)

func main() {
	// 1. ARITHMETIC OPERATORS
	// Standard math operations: + (add), - (subtract), * (multiply), / (divide), % (modulo/remainder)
	fmt.Println("--- Arithmetic Operators ---")
	a, b := 10, 3

	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)

	// Integer Division: When dividing two integers, Go truncates the decimal.
	// 10 / 3 evaluates to 3, not 3.333. 
	fmt.Println("Integer Division:", a/b) 

	// Modulo gives you the remainder of the division. (10 divided by 3 leaves a remainder of 1)
	fmt.Println("Modulo (Remainder):", a%b)

	// 2. COMPARISON OPERATORS
	// These operators always return a boolean (true or false).
	// == (equal), != (not equal), < (less), <= (less or equal), > (greater), >= (greater or equal)
	fmt.Println("\n--- Comparison Operators ---")
	x, y := 5, 10

	fmt.Println("x == y:", x == y)
	fmt.Println("x != y:", x != y)
	fmt.Println("x < y:", x < y)

	// 3. LOGICAL OPERATORS
	// Used to combine multiple boolean expressions.
	// && (Logical AND), || (Logical OR), ! (Logical NOT)
	fmt.Println("\n--- Logical Operators ---")
	isWeekend := true
	isSunny := false

	// AND: True only if BOTH sides are true
	fmt.Println("Go to beach (AND):", isWeekend && isSunny) 

	// OR: True if AT LEAST ONE side is true
	fmt.Println("Relax at home (OR):", isWeekend || isSunny) 

	// NOT: Inverts the boolean value
	fmt.Println("Is NOT sunny:", !isSunny) 

	// 4. BITWISE OPERATORS
	// These manipulate numbers at the raw binary level (1s and 0s).
	// & (AND), | (OR), ^ (XOR), << (Left Shift), >> (Right Shift)
	fmt.Println("\n--- Bitwise Operators ---")

	var bitA uint8 = 12 // Binary: 0000 1100
	var bitB uint8 = 10 // Binary: 0000 1010

	// Bitwise AND (&): 1 if both bits are 1.
	// Result: 0000 1000 (Decimal 8)
	fmt.Printf("Bitwise AND (bitA & bitB): %d\n", bitA&bitB)

	// Bitwise OR (|): 1 if at least one bit is 1.
	// Result: 0000 1110 (Decimal 14)
	fmt.Printf("Bitwise OR (bitA | bitB): %d\n", bitA|bitB)

	// Bitwise XOR (^): 1 if bits are different.
	// Result: 0000 0110 (Decimal 6)
	fmt.Printf("Bitwise XOR (bitA ^ bitB): %d\n", bitA^bitB)

	// 5. THE BIT CLEAR OPERATOR (&^)
	// This is unique to Go. It means "AND NOT". 
	// It clears the bits in the left operand wherever the right operand has a 1.
	// bitA: 0000 1100
	// bitB: 0000 1010
	// Res:  0000 0100 (Decimal 4) -> The '8' bit was cleared because bitB had it set.
	fmt.Printf("Bit Clear (bitA &^ bitB): %d\n", bitA&^bitB)
}
