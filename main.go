package main

import (
	"fmt"
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
	o.Backward()

	file, err := os.Create("graph.dot")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	/*_, err = file.WriteString(o.Graph().String())
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}*/

	mlp := NewMLP(3, []int{4, 4, 1})
	x := []*Value{NewValue(2, "x1"), NewValue(3, "x2"), NewValue(-1, "x3")}
	out := mlp.Forward(x)
	for _, v := range out {
		_, err = file.WriteString(v.Graph().String())
		if err != nil {
			log.Fatal("Error writing to file:", err)
		}
	}

	/*layer := NewNeuron(2, true)
	x := []*Value{NewValue(2, "x1"), NewValue(3, "x2")}
	out := layer.Forward(x)
	_, err = file.WriteString(out.Graph().String())
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}*/

	xs := [][]*Value{
		{NewValue(2, "x11"), NewValue(3, "x12"), NewValue(-1, "x13")},
		{NewValue(3, "x21"), NewValue(-1, "x22"), NewValue(0.5, "x23")},
		{NewValue(0.5, "x31"), NewValue(1, "x32"), NewValue(1, "x33")},
		{NewValue(1, "x41"), NewValue(1, "x42"), NewValue(-1, "x43")},
	}
	ys := []*Value{NewValue(1, "y1"), NewValue(-1, "y2"), NewValue(-1, "y3"), NewValue(1, "y4")}
	loss := NewValue(0, "loss")
	for i, x := range xs {
		p := mlp.Forward(x)
		diff := p[0].Sub(ys[i], "diff1")
		diff = diff.Mul(diff, "diff2")
		loss = loss.Add(diff, "loss")
	}
	fmt.Println(loss)
	loss.Backward()
}
