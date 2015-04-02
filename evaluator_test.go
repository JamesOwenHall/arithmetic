package main

import (
	"math"
	"strings"
	"testing"
)

func TestEvaluate(t *testing.T) {
	// Single number
	input := "4"
	result, err := Evaluate(strings.NewReader(input))
	expected := 4.0
	if err != nil {
		t.Error(err)
	} else if !roughly(result, expected) {
		t.Errorf("Expected %f, got %f", expected, result)
	}

	// Empty program
	input = ""
	_, err = Evaluate(strings.NewReader(input))
	if err != ErrEmptyProgram {
		t.Errorf("Expected empty program error, got %v", err)
	}

	// Addition
	input = "(+ 1 2)"
	result, err = Evaluate(strings.NewReader(input))
	expected = 3.0
	if err != nil {
		t.Error(err)
	} else if !roughly(result, expected) {
		t.Errorf("Expected %f, got %f", expected, result)
	}

	// Subtraction
	input = "(- 1 2.2)"
	result, err = Evaluate(strings.NewReader(input))
	expected = -1.2
	if err != nil {
		t.Error(err)
	} else if !roughly(result, expected) {
		t.Errorf("Expected %f, got %f", expected, result)
	}

	// Multiplication
	input = "(* -4 -2.5)"
	result, err = Evaluate(strings.NewReader(input))
	expected = 10
	if err != nil {
		t.Error(err)
	} else if !roughly(result, expected) {
		t.Errorf("Expected %f, got %f", expected, result)
	}

	// Division
	input = "(/ 5 2)"
	result, err = Evaluate(strings.NewReader(input))
	expected = 2.5
	if err != nil {
		t.Error(err)
	} else if !roughly(result, expected) {
		t.Errorf("Expected %f, got %f", expected, result)
	}

	// Nesting
	input = "(+ (/ 27 3 3) (* 3.14 (- 2 1.1)))"
	result, err = Evaluate(strings.NewReader(input))
	expected = 5.826
	if err != nil {
		t.Error(err)
	} else if !roughly(result, expected) {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func roughly(a, b float64) bool {
	return math.Abs(b-a) < 0.0001
}
