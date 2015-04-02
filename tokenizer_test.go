package main

import (
	"strings"
	"testing"
)

func TestTokenize(t *testing.T) {
	// Empty string
	input := ""
	c := Tokenize(strings.NewReader(input))
	expected := []Token{}
	if ok, want, got := ExpectTokens(expected, c); !ok {
		t.Errorf("Expected token %v, got %v", want, got)
	}

	// Parentheses
	input = "()"
	c = Tokenize(strings.NewReader(input))
	expected = []Token{
		{Type: OpenParen, Val: "("},
		{Type: CloseParen, Val: ")"},
	}
	if ok, want, got := ExpectTokens(expected, c); !ok {
		t.Errorf("Expected token %v, got %v", want, got)
	}

	// Whitespace
	input = " \t(\t \n \r )    "
	c = Tokenize(strings.NewReader(input))
	expected = []Token{
		{Type: OpenParen, Val: "("},
		{Type: CloseParen, Val: ")"},
	}
	if ok, want, got := ExpectTokens(expected, c); !ok {
		t.Errorf("Expected token %v, got %v", want, got)
	}

	// Operators
	input = "(+ - / *)"
	c = Tokenize(strings.NewReader(input))
	expected = []Token{
		{Type: OpenParen, Val: "("},
		{Type: Operator, Val: "+"},
		{Type: Operator, Val: "-"},
		{Type: Operator, Val: "/"},
		{Type: Operator, Val: "*"},
		{Type: CloseParen, Val: ")"},
	}
	if ok, want, got := ExpectTokens(expected, c); !ok {
		t.Errorf("Expected token %v, got %v", want, got)
	}

	// Numbers
	input = "(+ 2.23 - -0.1 9 / *)"
	c = Tokenize(strings.NewReader(input))
	expected = []Token{
		{Type: OpenParen, Val: "("},
		{Type: Operator, Val: "+"},
		{Type: Number, Val: "2.23"},
		{Type: Operator, Val: "-"},
		{Type: Number, Val: "-0.1"},
		{Type: Number, Val: "9"},
		{Type: Operator, Val: "/"},
		{Type: Operator, Val: "*"},
		{Type: CloseParen, Val: ")"},
	}
	if ok, want, got := ExpectTokens(expected, c); !ok {
		t.Errorf("Expected token %v, got %v", want, got)
	}
}

// ExpectTokens returns true if the Tokens produced by c is equal to the
// "tokens" argument.
func ExpectTokens(tokens []Token, c <-chan Token) (bool, *Token, *Token) {
	var i int
	for token := range c {
		if i >= len(tokens) {
			return false, nil, &token
		}

		if tokens[i] != token {
			return false, &tokens[i], &token
		}

		i++
	}

	if i != len(tokens) {
		return false, &tokens[i], nil
	}

	return true, nil, nil
}
