package strategy

import "testing"

func TestStrategy(t *testing.T) {
	operator := Operator{}

	operator.setStrategy(&add{})
	testAddSamples := []struct {
		input [2]int
		want  int
	}{
		{[2]int{1, 2}, 3},
		{[2]int{3, 5}, 8},
	}

	for _, tt := range testAddSamples {
		if got := operator.calculate(tt.input[0], tt.input[1]); tt.want != got {
			t.Errorf("Add() = %v, want %v", got, tt.want)
		}
	}

	operator.setStrategy(&reduce{})
	testReduceSamples := []struct {
		input [2]int
		want  int
	}{
		{[2]int{1, 2}, -1},
		{[2]int{3, 5}, -2},
	}
	for _, tt := range testReduceSamples {
		if got := operator.calculate(tt.input[0], tt.input[1]); tt.want != got {
			t.Errorf("Reduce() = %v, want %v", got, tt.want)
		}
	}
}
