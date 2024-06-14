package main

import "fmt"

func main() {
	// Create a new Value
	v := Value{data: 1.0}
	// Create another Value
	v2 := Value{data: 2.0}
	// Add the two Values together
	result := v.Add(v2).Mul(v2)
	// Print the result
	fmt.Println(result)
}
