package math_examples

import (
	"testing"
)

func TestGetAngle(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"12:00", "Least angle for 12:00 is 0."},
		{"3:00", "Least angle for 3:00 is 90."},
		{"6:00", "Least angle for 6:00 is 180."},
		{"9:00", "Least angle for 9:00 is 90."},
		{"1:15", "Least angle for 1:15 is 52.5."},
		{"13:00", "Least angle for 13:00 is 30."},
		{"18:00", "Least angle for 18:00 is 180."},
	}

	for _, test := range tests {
		got, err := GetAngle(test.input)
		if err != nil {
			t.Errorf("GetAngle(%q) returned an error: %v", test.input, err)
		}
		if got != test.want {
			t.Errorf("GetAngle(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

func TestGetAngleError(t *testing.T) {
	tests := []string{
		"25:00", // hour out of bounds
		"12:75", // minute out of bounds
		"123",   // not a valid time string
		"abc",   // not a valid time string
		"24:01", // minute out-of-bounds in 24-hour format
	}

	for _, test := range tests {
		_, err := GetAngle(test)
		if err == nil {
			t.Errorf("GetAngle(%q) should have returned an error", test)
		}
	}
}
