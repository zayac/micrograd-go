package main

import (
	"reflect"
	"testing"
)

func TestValueAdd(t *testing.T) {
	tests := []struct {
		name string
		v1   Value
		v2   Value
		want Value
	}{{
		name: "positive values",
		v1:   Value{data: 5.0},
		v2:   Value{data: 3.0},
		want: Value{
			data: 8.0,
			prev: []*Value{{data: 5.0}, {data: 3.0}},
			op:   AddOp,
		},
	}, {
		name: "negative values",
		v1:   Value{data: -5.0},
		v2:   Value{data: -3.0},
		want: Value{
			data: -8.0,
			prev: []*Value{{data: -5.0}, {data: -3.0}},
			op:   AddOp,
		},
	}, {
		name: "zero values",
		v1:   Value{data: 0.0},
		v2:   Value{data: 0.0},
		want: Value{
			data: 0.0,
			prev: []*Value{{data: 0.0}, {data: 0.0}},
			op:   AddOp,
		},
	},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v1.Add(tt.v2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, but got %v", tt.want, got)
			}
		})
	}
}

func TestValueMul(t *testing.T) {
	tests := []struct {
		name         string
		v1, v2, want Value
	}{
		{
			name: "positive values",
			v1:   Value{data: 5.0},
			v2:   Value{data: 3.0},
			want: Value{
				data: 15.0,
				prev: []*Value{{data: 5.0}, {data: 3.0}},
				op:   MulOp,
			},
		},
		{
			name: "negative values",
			v1:   Value{data: -5.0},
			v2:   Value{data: -3.0},
			want: Value{
				data: 15.0,
				prev: []*Value{{data: -5.0}, {data: -3.0}},
				op:   MulOp,
			},
		},
		{
			name: "zero values",
			v1:   Value{data: 0.0},
			v2:   Value{data: 0.0},
			want: Value{
				data: 0.0,
				prev: []*Value{{data: 0.0}, {data: 0.0}},
				op:   MulOp,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v1.Mul(tt.v2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, but got %v", tt.want, got)
			}
		})
	}
}
