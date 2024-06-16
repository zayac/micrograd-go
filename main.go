package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	v := NewValue(5.0, "a")
	v2 := NewValue(3.0, "b")
	result := v.Add(v2, "c").Mul(v2, "d")
	fmt.Println(result)

	file, err := os.Create("graph.dot")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	_, err = file.WriteString(result.Graph().String())
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}
}
