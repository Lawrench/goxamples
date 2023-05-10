package math_examples

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"Input is -85", -85, false},
		{"Input is 1", 1, false},
		{"Input is 2", 2, true},
		{"Input is 4", 4, false},
		{"Input is 17", 17, true},
		{"Input is 121", 121, false},
		{"Input is 11", 11, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsPrime(test.input)
			if got != test.want {
				t.Errorf("IsPrime(%d) = %v, want %v", test.input, got, test.want)
			}
		})
	}
}
