package main

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrEmptyProgram = errors.New("The source program is empty.")
)

type ErrUnknownOperator Node

func (e ErrUnknownOperator) Error() string {
	return "Unknown operator " + e.Text
}

func Evaluate(r io.Reader) (float64, error) {
	root, err := Parse(r)
	if err != nil {
		return 0, err
	}

	if len(root.Children) == 0 {
		return 0, ErrEmptyProgram
	}

	return evaluateNode(root.Children[0])
}

func evaluateNode(n Node) (float64, error) {
	if n.Type == NodeNumber {
		var result float64
		_, err := fmt.Sscanf(n.Text, "%f", &result)
		return result, err
	}

	switch n.Text {
	case "+":
		var result float64
		for _, child := range n.Children {
			childResult, err := evaluateNode(child)
			if err != nil {
				return 0, err
			}
			result += childResult
		}
		return result, nil
	case "-":
		if len(n.Children) == 0 {
			return 0, nil
		}

		result, err := evaluateNode(n.Children[0])
		if err != nil {
			return 0, err
		}

		for i := 1; i < len(n.Children); i++ {
			childResult, err := evaluateNode(n.Children[i])
			if err != nil {
				return 0, err
			}
			result -= childResult
		}

		return result, nil
	case "*":
		var result float64 = 1.0
		for _, child := range n.Children {
			childResult, err := evaluateNode(child)
			if err != nil {
				return 0, err
			}
			result *= childResult
		}
		return result, nil
	case "/":
		if len(n.Children) == 0 {
			return 1, nil
		}

		result, err := evaluateNode(n.Children[0])
		if err != nil {
			return 0, err
		}

		for i := 1; i < len(n.Children); i++ {
			childResult, err := evaluateNode(n.Children[i])
			if err != nil {
				return 0, err
			}
			result /= childResult
		}

		return result, nil
	default:
		return 0, ErrUnknownOperator(n)
	}
}
