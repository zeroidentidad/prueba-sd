package main

import (
	"reflect"
	"testing"
)

// go test -run TestSumToZero -v
func TestSumToZero(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Base case",
			input:    []int{3, 4, -7, 5, -6, 2, 5, -1, 8},
			expected: []int{5, 8},
		},
		{
			name:     "Arreglo vacío",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Arreglo nulo",
			input:    nil,
			expected: []int{},
		},
		{
			name:     "Todos positivos",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Todos negativos",
			input:    []int{-1, -2, -3, -4, -5},
			expected: []int{-1, -2, -3, -4, -5},
		},
		{
			name:     "Con secuencia de suma cero",
			input:    []int{1, 2, -3, 1, 2, -3},
			expected: []int{},
		},
		{
			name:     "Con sumas cero juntas",
			input:    []int{1, -1, 2, -2, 3, -3},
			expected: []int{},
		},
		{
			name:     "Secuencia al inicio suma cero",
			input:    []int{-3, 3, 1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Secuencia al final suma cero",
			input:    []int{1, 2, 3, -3, -2, -1},
			expected: []int{},
		},
		{
			name:     "Secuencia que suma cero repetida en la mitad",
			input:    []int{1, 2, -3, 3, 4, -4, 5},
			expected: []int{3, 5},
		},
		{
			name:     "Arreglo con solo un elemento",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Arreglo con un único cero",
			input:    []int{0},
			expected: []int{0},
		},
	}

	// Run table cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumToZero(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("failed %s: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}
