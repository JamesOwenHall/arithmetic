package main

import (
	"errors"
	"fmt"
	"io"
)

// NodeType is the type associated to a node.
type NodeType int8

const (
	NodeOperation NodeType = iota
	NodeNumber
)

var (
	ErrEndOfFile = errors.New("Unexpected end of file.")
)

type ErrUnexpectedToken Token

func (e ErrUnexpectedToken) Error() string {
	return fmt.Sprintf("Unexpected token %s.", e.Val)
}

// Node is a single node within the abstract syntax tree.
type Node struct {
	Type     NodeType
	Children []Node
	Text     string
}

// Parse returns a root node representing the AST.
func Parse(r io.Reader) (Node, error) {
	c := Tokenize(r)
	root := Node{
		Children: make([]Node, 0),
		Text:     "",
	}

	for {
		n, err := parseNode(c)
		if err != nil {
			return root, err
		}

		if n != nil {
			root.Children = append(root.Children, *n)
		} else {
			break
		}
	}

	return root, nil
}

func parseNode(c <-chan Token) (*Node, error) {
	t, ok := <-c
	if !ok {
		return nil, nil
	}

	switch t.Type {
	case TokNumber:
		return numberNode(t), nil
	case TokOpenParen:
		return operationNode(c)
	default:
		return nil, ErrUnexpectedToken(t)
	}
}

func numberNode(t Token) *Node {
	return &Node{
		Type: NodeNumber,
		Text: t.Val,
	}
}

func operationNode(c <-chan Token) (*Node, error) {
	n := new(Node)

	next, ok := <-c
	if !ok {
		return nil, ErrEndOfFile
	}

	if next.Type != TokOperator {
		return nil, ErrUnexpectedToken(next)
	}

	n.Text = next.Val

loop:
	for {
		next, ok := <-c
		if !ok {
			return nil, ErrEndOfFile
		}

		switch next.Type {
		case TokOperator:
			return nil, ErrUnexpectedToken(next)
		case TokNumber:
			n.Children = append(n.Children, *numberNode(next))
		case TokCloseParen:
			break loop
		default:
			child, err := operationNode(c)
			if err != nil {
				return nil, err
			}

			n.Children = append(n.Children, *child)
		}
	}

	return n, nil
}
