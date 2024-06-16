package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	v := Value{data: 1.0}
	v2 := Value{data: 2.0}
	result := v.Add(v2).Mul(v2)
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
