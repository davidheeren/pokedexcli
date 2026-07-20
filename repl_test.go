package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "ONE two ThrEe",
			expected: []string{"one", "two", "three"},
		},
		{
			input:    "this has a\nnewline",
			expected: []string{"this", "has", "a", "newline"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if !reflect.DeepEqual(c.expected, actual) {
			t.Fatalf("expected: %v, got: %v", c.expected, actual)
		}
	}
}
