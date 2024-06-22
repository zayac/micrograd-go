package main

import (
	"fmt"
	"math/rand/v2"
)

type Neuron struct {
	w        []*Value
	b        *Value
	isLinear bool
}

func NewID() string {
	n := rand.Int32()
	return fmt.Sprintf("%d", n)
}

func NewNeuron(nin int, isLinear bool) *Neuron {
	ret := &Neuron{
		w:        make([]*Value, nin),
		b:        NewValue(rand.Float64()*2-1, NewID()),
		isLinear: isLinear,
	}
	for i := range ret.w {
		ret.w[i] = NewValue(rand.Float64()*2-1, NewID())
	}
	return ret
}

func (n *Neuron) Forward(x []*Value) *Value {
	sum := n.b
	for i := range x {
		sum = sum.Add(n.w[i].Mul(x[i], NewID()), NewID())
	}
	if !n.isLinear {
		return sum.Relu(fmt.Sprintf("relu-%s", NewID()))
	}
	return sum
}

type Layer struct {
	neurons []*Neuron
}

func NewLayer(nin, nout int, isLinear bool) *Layer {
	ret := Layer{
		neurons: make([]*Neuron, nout),
	}
	for i := range ret.neurons {
		ret.neurons[i] = NewNeuron(nin, isLinear)
	}
	return &ret
}

func (l *Layer) Forward(x []*Value) []*Value {
	ret := make([]*Value, len(l.neurons))
	for i := range l.neurons {
		ret[i] = l.neurons[i].Forward(x)
	}
	return ret
}

type MLP struct {
	layers []*Layer
}

func NewMLP(nin int, nouts []int) *MLP {
	sz := append([]int{nin}, nouts...)
	layers := make([]*Layer, len(nouts))
	for i := 0; i < len(nouts); i++ {
		isLinear := i != len(nouts)-1
		layers[i] = NewLayer(sz[i], sz[i+1], isLinear)
	}

	return &MLP{
		layers: layers,
	}
}

func (m *MLP) Forward(x []*Value) []*Value {
	for i := range m.layers {
		x = m.layers[i].Forward(x)
	}
	return x
}
