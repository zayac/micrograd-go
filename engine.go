package main

import (
	"fmt"
	"log"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type op int

const (
	UnknownOp op = iota
	AddOp
	MulOp
)

type Value struct {
	data float64
	prev []*Value
	op   op
}

func (v Value) String() string {
	return fmt.Sprintf("%.2f", v.data)
}

func (v Value) Add(v2 Value) Value {
	return Value{
		data: v.data + v2.data,
		prev: []*Value{&v, &v2},
		op:   AddOp,
	}
}

func (v Value) Mul(v2 Value) Value {
	return Value{
		data: v.data * v2.data,
		prev: []*Value{&v, &v2},
		op:   MulOp,
	}
}

func (v Value) buildGraph(graph *cgraph.Graph) (*cgraph.Node, error) {
	node, err := graph.CreateNode(v.String())
	if err != nil {
		return nil, err
	}
	for _, p := range v.prev {
		prevNode, err := p.buildGraph(graph)
		if err != nil {
			return nil, err
		}
		_, err = graph.CreateEdge("", prevNode, node)
		if err != nil {
			return nil, err
		}
	}
	return node, nil
}

func Graph(g *graphviz.Graphviz, v Value) (*cgraph.Graph, error) {
	graph, err := g.Graph()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := g.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if _, err = v.buildGraph(graph); err != nil {
		return nil, err
	}
	return graph, nil
}
