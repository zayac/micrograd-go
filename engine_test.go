package main

import (
	"reflect"
	"testing"
)

func TestValueAdd(t *testing.T) {
	tests := []struct {
		name string
		v1   *Value
		v2   *Value
		want *Value
	}{{
		name: "positive values",
		v1:   NewValue(5.0, "a"),
		v2:   NewValue(3.0, "b"),
		want: &Value{
			data: 8.0,
			prev: []*Value{
				{data: 5.0, label: "a"},
				{data: 3.0, label: "b"},
			},
			op:    AddOp,
			label: "c",
		},
	}, {
		name: "negative values",
		v1:   NewValue(-5.0, "a"),
		v2:   NewValue(-3.0, "b"),
		want: &Value{
			data: -8.0,
			prev: []*Value{
				{data: -5.0, label: "a"},
				{data: -3.0, label: "b"},
			},
			op:    AddOp,
			label: "c",
		},
	}, {
		name: "zero values",
		v1:   NewValue(0.0, "a"),
		v2:   NewValue(0.0, "b"),
		want: &Value{
			data: 0.0,
			prev: []*Value{
				{data: 0.0, label: "a"},
				{data: 0.0, label: "b"},
			},
			op:    AddOp,
			label: "c",
		},
	},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v1.Add(tt.v2, "c")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %+v, but got %+v", tt.want, got)
			}
		})
	}
}

func TestValueMul(t *testing.T) {
	tests := []struct {
		name         string
		v1, v2, want *Value
	}{
		{
			name: "positive values",
			v1:   NewValue(5.0, "a"),
			v2:   NewValue(3.0, "b"),
			want: &Value{
				data: 15.0,
				prev: []*Value{
					{data: 5.0, label: "a"},
					{data: 3.0, label: "b"},
				},
				op:    MulOp,
				label: "c",
			},
		},
		{
			name: "negative values",
			v1:   NewValue(-5.0, "a"),
			v2:   NewValue(-3.0, "b"),
			want: &Value{
				data: 15.0,
				prev: []*Value{
					{data: -5.0, label: "a"},
					{data: -3.0, label: "b"},
				},
				op:    MulOp,
				label: "c",
			},
		},
		{
			name: "zero values",
			v1:   NewValue(0.0, "a"),
			v2:   NewValue(0.0, "b"),
			want: &Value{
				data: 0.0,
				prev: []*Value{
					{data: 0.0, label: "a"},
					{data: 0.0, label: "b"},
				},
				op:    MulOp,
				label: "c",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v1.Mul(tt.v2, "c")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %+v, but got %+v", tt.want, got)
			}
		})
	}
}
