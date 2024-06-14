package main

import "fmt"

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
