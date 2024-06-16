package main

import (
	"fmt"

	"github.com/emicklei/dot"
)

type op int

const (
	UnknownOp op = iota
	AddOp
	MulOp
)

func (o op) String() string {
	switch o {
	case AddOp:
		return "+"
	case MulOp:
		return "*"
	default:
		panic("unknown operation")
	}
}

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

func (v Value) buildGraph(g *dot.Graph) dot.Node {
	node := g.Node(v.String()).Box()
	var op dot.Node
	if v.op != UnknownOp {
		op = g.Node(v.op.String())
		g.Edge(op, node)
	}
	for _, p := range v.prev {
		prevNode := p.buildGraph(g)
		if v.op == UnknownOp {
			g.Edge(prevNode, node)
		} else {
			g.Edge(prevNode, op)
		}
	}
	return node
}

func (v Value) Graph() *dot.Graph {
	g := dot.NewGraph(dot.Directed)
	g.Attr("rankdir", "LR")
	v.buildGraph(g)
	return g
}
