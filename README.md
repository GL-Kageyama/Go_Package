# Go Package

## Package

## Code
```Go
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
```
```Go
package calculation

func CalcAdd(x int, y int) int {
	result := x + y
	return result
}
```
```Go
package calculation

func CalcSub(x int, y int) int {
	result := x - y
	return result
}
```
```Go
package calculation

func CalcMulti(x int, y int) int {
	result := x * y
	return result
}
```
```Go
package calculation

func CalcDiv(x int, y int) int {
	result := x / y
	return result
}
```

## Output Sample
~ $ go build -o calc main.go  
~ $ ./calc  
Addition       : 15  
Subtraction    : 5  
Multiplication : 50  
Division       : 2  

## Note
For go1.13 or later, the environment variable GO111MODULE must be set to "auto".  
$ go env -w GO111MODULE=auto  
