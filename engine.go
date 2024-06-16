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
	data, grad float64
	prev       []*Value
	op         op
	label      string
}

func NewValue(data float64, label string) Value {
	return Value{data: data, label: label}
}

func (v Value) Add(v2 Value, label string) Value {
	return Value{
		data:  v.data + v2.data,
		prev:  []*Value{&v, &v2},
		op:    AddOp,
		label: label,
	}
}

func (v Value) Mul(v2 Value, label string) Value {
	return Value{
		data:  v.data * v2.data,
		prev:  []*Value{&v, &v2},
		op:    MulOp,
		label: label,
	}
}

func (v Value) buildGraph(g *dot.Graph) dot.Node {
	node := g.Node(v.label).NewRecordBuilder()
	node.FieldWithId(v.label, "label")
	node.FieldWithId(fmt.Sprintf("data %.2f", v.data), "data")
	node.FieldWithId(fmt.Sprintf("grad %.2f", v.grad), "grad")
	node.Build()
	var op dot.Node
	if v.op != UnknownOp {
		op = g.Node(v.op.String())
		g.Edge(op, g.Node(v.label))
	}
	for _, p := range v.prev {
		prevNode := p.buildGraph(g)
		if v.op == UnknownOp {
			g.Edge(prevNode, g.Node(v.label))
		} else {
			g.Edge(prevNode, op)
		}
	}
	return g.Node(v.label)
}

func (v Value) Graph() *dot.Graph {
	g := dot.NewGraph(dot.Directed)
	g.Attr("rankdir", "LR")
	v.buildGraph(g)
	return g
}
