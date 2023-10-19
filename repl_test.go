package main

import "testing"

func TestCleanImput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Error("the length are not equal: %w", len(actual), len(cs.expected))
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Error("%v does not equil %v", actualWord, expectedWord)
			}
		}

	}
}
