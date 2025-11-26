package main

import(
	"testing"
	"fmt"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input		string
		expected	[]string
	}{

			{
				input: "  hello world  ",
				expected: []string{"hello", "world"},
			},
			{
				input: "Charmander Bulbasaur PIKACHU",
				expected: []string{"charmander", "bulbasaur", "pikachu"},
			},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Printf("input: %s, output: %v, expected: %v",c.input, actual, c.expected)
		if len(actual) != len(c.expected) {
			t.Errorf("Lenght expected :%d, got %d", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word expected: %s, got %s", expectedWord, word)
			}
		}
	}
}
