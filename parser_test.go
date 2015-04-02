package main

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	// Empty string
	input := ""
	ast, err := Parse(strings.NewReader(input))
	expected := Node{
		Children: nil,
	}
	if err != nil {
		t.Error(err)
	} else if ok, want, got := ExpectAst(expected, ast); !ok {
		t.Errorf("Expected node %v, got %v", want, got)
	}

	// Number
	input = "-9.2"
	ast, err = Parse(strings.NewReader(input))
	expected = Node{
		Children: []Node{
			{Type: NodeNumber, Text: "-9.2"},
		},
	}
	if err != nil {
		t.Error(err)
	} else if ok, want, got := ExpectAst(expected, ast); !ok {
		t.Errorf("Expected node %v, got %v", want, got)
	}

	// Operation
	input = "(+ 1 -2 3.4 -5.6)"
	ast, err = Parse(strings.NewReader(input))
	expected = Node{
		Children: []Node{
			{
				Type: NodeOperation,
				Text: "+",
				Children: []Node{
					{Type: NodeNumber, Text: "1"},
					{Type: NodeNumber, Text: "-2"},
					{Type: NodeNumber, Text: "3.4"},
					{Type: NodeNumber, Text: "-5.6"},
				},
			},
		},
	}
	if err != nil {
		t.Error(err)
	} else if ok, want, got := ExpectAst(expected, ast); !ok {
		t.Errorf("Expected node %v, got %v", want, got)
	}

	// Nested operations
	input = "(+ (* 2 3) 4)"
	ast, err = Parse(strings.NewReader(input))
	expected = Node{
		Children: []Node{
			{
				Type: NodeOperation,
				Text: "+",
				Children: []Node{
					{
						Type: NodeOperation,
						Text: "*",
						Children: []Node{
							{Type: NodeNumber, Text: "2"},
							{Type: NodeNumber, Text: "3"},
						},
					},
					{Type: NodeNumber, Text: "4"},
				},
			},
		},
	}
	if err != nil {
		t.Error(err)
	} else if ok, want, got := ExpectAst(expected, ast); !ok {
		t.Errorf("Expected node %v, got %v", want, got)
	}
}

func ExpectAst(want, got Node) (bool, *Node, *Node) {
	if want.Type != got.Type || want.Text != got.Text {
		return false, &want, &got
	}

	if len(want.Children) != len(got.Children) {
		return false, &want, &got
	}

	for i := range want.Children {
		if ok, want1, got1 := ExpectAst(want.Children[i], got.Children[i]); !ok {
			return false, want1, got1
		}
	}

	return true, nil, nil
}
