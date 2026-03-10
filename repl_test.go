package main

import "testing"

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
			input:    "",
			expected: []string{},
		},
		{
			input:    "single",
			expected: []string{"single"},
		},
		{
			input:    " multiple   spaces here ",
			expected: []string{"multiple", "spaces", "here"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) length = %d; want %d; actual=%v", c.input, len(actual), len(c.expected), actual)
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%q)[%d] = %q; want %q", c.input, i, actual[i], c.expected[i])
			}
		}
	}
}
