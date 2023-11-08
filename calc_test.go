package devopscrashcourse

import (
	"testing"
)

type args struct {
	a int
	b int
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 + 1 = 2", args{1, 1}, 2},
		{"1 + -1 = 0", args{1, -1}, 0},
		{"0 + 0 = 0", args{0, 0}, 0},
		{"-1 + -1 = -2", args{-1, -1}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 - 1 = 0", args{1, 1}, 0},
		{"1 - -1 = 2", args{1, -1}, 2},
		{"0 - 0 = 0", args{0, 0}, 0},
		{"-1 - -1 = 0", args{-1, -1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 * 2 = 1", args{1, 2}, 2},
		{"1 * -1 = -1", args{1, -1}, -1},
		{"0 * 0 = 0", args{0, 0}, 0},
		{"-1 * -1 = 1", args{-1, -1}, 1},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 / 2 = 1", args{1, 2}, 0},
		{"1 / -1 = -1", args{1, -1}, -1},
		{"0 / 1 = 0", args{0, 1}, 0},
		{"-1 / -1 = 1", args{-1, -1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divide(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModulo(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 % 2 = 0", args{1, 2}, 1},
		{"1 % -1 = 0", args{1, -1}, 0},
		{"0 % 1 = 0", args{0, 1}, 0},
		{"-1 % -1 = 0", args{-1, -1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Modulo(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Modulo() = %v, want %v", got, tt.want)
			}
		})
	}
}
