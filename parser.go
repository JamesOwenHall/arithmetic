package main

import (
	"fmt"
	"io"
)

type NodeType int8

const (
	NodeOperation NodeType = iota
	NodeNumber
)

var (
	ErrEndOfFile = fmt.Errorf("Unexpected end of file.")
)

type ErrUnexpectedToken Token

func (e ErrUnexpectedToken) Error() string {
	return fmt.Sprintf("Unexpected token %s.", e.Val)
}

type Node struct {
	Type     NodeType
	Children []Node
	Text     string
}

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
		return NumberNode(t), nil
	case TokOpenParen:
		return OperationNode(c)
	default:
		return nil, ErrUnexpectedToken(t)
	}
}

func NumberNode(t Token) *Node {
	return &Node{
		Type: NodeNumber,
		Text: t.Val,
	}
}

func OperationNode(c <-chan Token) (*Node, error) {
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
			n.Children = append(n.Children, *NumberNode(next))
		case TokCloseParen:
			break loop
		default:
			child, err := OperationNode(c)
			if err != nil {
				return nil, err
			}

			n.Children = append(n.Children, *child)
		}
	}

	return n, nil
}
