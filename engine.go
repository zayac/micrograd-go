package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/emicklei/dot"
)

type op int

const (
	UnknownOp op = iota
	AddOp
	SubOp
	MulOp
	TanhOp
	ReluOp
)

func (o op) String() string {
	switch o {
	case AddOp:
		return "+"
	case SubOp:
		return "-"
	case MulOp:
		return "*"
	case TanhOp:
		return "tanh"
	case ReluOp:
		return "relu"
	default:
		panic("unknown operation")
	}
}

type Value struct {
	data, grad float64
	prev       []*Value
	op         op
	label      string
	backward   func()
}

func NewValue(data float64, label string) *Value {
	return &Value{data: data, label: label, backward: func() {}}
}

func (v *Value) Add(v2 *Value, label string) *Value {
	ret := &Value{
		data:  v.data + v2.data,
		prev:  []*Value{v, v2},
		op:    AddOp,
		label: label,
	}
	ret.backward = func() {
		v.grad += ret.grad
		v2.grad += ret.grad
	}
	return ret
}

func (v *Value) Sub(v2 *Value, label string) *Value {
	ret := &Value{
		data:  v.data + v2.data,
		prev:  []*Value{v, v2},
		op:    SubOp,
		label: label,
	}
	ret.backward = func() {
		v.grad -= ret.grad
		v2.grad -= ret.grad
	}
	return ret
}

func (v *Value) Mul(v2 *Value, label string) *Value {
	ret := &Value{
		data:  v.data * v2.data,
		prev:  []*Value{v, v2},
		op:    MulOp,
		label: label,
	}
	ret.backward = func() {
		v.grad += v2.data * ret.grad
		v2.grad += v.data * ret.grad
	}
	return ret
}

func (v *Value) Tanh(label string) *Value {
	x := v.data
	ret := &Value{
		data:  (math.Exp(2*x) - 1) / (math.Exp(2*x) + 1),
		prev:  []*Value{v},
		op:    TanhOp,
		label: label,
	}
	ret.backward = func() {
		v.grad += (1 - ret.data*ret.data) * ret.grad
	}
	return ret
}

func (v *Value) Relu(label string) *Value {
	var data float64
	if v.data > 0 {
		data = v.data
	}
	ret := &Value{
		data:  data,
		prev:  []*Value{v},
		op:    ReluOp,
		label: label,
	}
	ret.backward = func() {
		if ret.data > 0 {
			v.grad += ret.grad
		}
	}
	return ret
}

func (v Value) buildGraph(g *dot.Graph) dot.Node {
	node := g.Node(v.label).NewRecordBuilder()
	node.FieldWithId(v.label, "label")
	node.FieldWithId(fmt.Sprintf("data %.2f", v.data), "data")
	node.FieldWithId(fmt.Sprintf("grad %.2f", v.grad), "grad")
	node.Build()
	var op dot.Node
	if v.op != UnknownOp {
		op = g.Node(v.label + "_op").Label(v.op.String())
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

func (v *Value) Backward() {
	seen := make(map[*Value]bool)
	var order []*Value
	var visitAll func(*Value)
	visitAll = func(val *Value) {
		if !seen[val] {
			seen[val] = true
			for _, p := range val.prev {
				visitAll(p)
			}
			order = append(order, val)
		}
	}
	visitAll(v)
	slices.Reverse(order)

	v.grad = 1
	for _, val := range order {
		val.backward()
	}
}
