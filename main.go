package main

import (
	"fmt"
	"./calculation"
)

// Main function
func main() {

	// Calculation with functions in the calculation package
	add   := calculation.CalcAdd(10, 5)
	sub   := calculation.CalcSub(10, 5)
	multi := calculation.CalcMulti(10, 5)
	div   := calculation.CalcDiv(10, 5)

	// Output Value
	fmt.Printf("Addition       : %d \n", add)
	fmt.Printf("Subtraction    : %d \n", sub)
	fmt.Printf("Multiplication : %d \n", multi)
	fmt.Printf("Division       : %d \n", div)
}

// =================================
//           Output Sample
// =================================
// ~ $ go build -o calc main.go
// ~ $ ./calc
// Addition       : 15
// Subtraction    : 5
// Multiplication : 50
// Division       : 2
