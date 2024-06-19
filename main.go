package main

import (
	"log"
	"os"
)

func main() {
	x1 := NewValue(2, "x1")
	x2 := NewValue(0, "x2")
	w1 := NewValue(-3, "w1")
	w2 := NewValue(1, "w2")
	b := NewValue(6.8813735870195432, "b")
	x1w1 := x1.Mul(w1, "x1*w1")
	x2w2 := x2.Mul(w2, "x2*w2")
	x1w1x2w2 := x1w1.Add(x2w2, "x1w1+x2w2")
	n := x1w1x2w2.Add(b, "n")
	o := n.Tanh("o")
	o.grad = 1
	o.backward()
	n.backward()
	b.backward()
	x1w1x2w2.backward()
	x1w1.backward()
	x2w2.backward()

	file, err := os.Create("graph.dot")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	_, err = file.WriteString(o.Graph().String())
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}
}
