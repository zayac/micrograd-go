package main

import (
	"fmt"
	"log"

	"github.com/goccy/go-graphviz"
)

func main() {
	v := Value{data: 1.0}
	v2 := Value{data: 2.0}
	result := v.Add(v2).Mul(v2)
	fmt.Println(result)

	g := graphviz.New()
	graph, err := Graph(g, result)
	if err != nil {
		log.Fatal(err)
	}
	if err := g.RenderFilename(graph, graphviz.PNG, "graph.png"); err != nil {
		log.Fatal("graphviz graph rendering failed: ", err)
	}
}
