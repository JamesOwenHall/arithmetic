package main

import (
	"bufio"
	"bytes"
	"io"
)

// TokenType is the type associated to a token.
type TokenType int8

const (
	TokError TokenType = iota
	TokOpenParen
	TokCloseParen
	TokOperator
	TokNumber
)

// Token is a keyword, number or parenthesis.
type Token struct {
	Type TokenType
	Val  string
}

// Tokenize returns a read-only channel that produces tokens from r.
func Tokenize(r io.Reader) <-chan Token {
	result := make(chan Token)

	go func(r io.Reader, out chan Token) {
		reader := bufio.NewReaderSize(r, 4)

	loop:
		for {
			next, err := peekRune(reader)
			if err != nil {
				break loop
			}

			switch next {
			case ' ':
				fallthrough
			case '\t':
				fallthrough
			case '\r':
				fallthrough
			case '\n':
				reader.ReadByte()
			case '(':
				reader.ReadByte()
				out <- Token{
					Type: TokOpenParen,
					Val:  "(",
				}
			case ')':
				reader.ReadByte()
				out <- Token{
					Type: TokCloseParen,
					Val:  ")",
				}
			default:
				word := getWord(reader)
				if len(word) == 0 {
					break loop
				}

				if word == "+" || word == "-" || word == "*" || word == "/" {
					out <- Token{
						Type: TokOperator,
						Val:  word,
					}
				} else {
					out <- Token{
						Type: TokNumber,
						Val:  word,
					}
				}
			}
		}

		close(out)
	}(r, result)

	return result
}

// getWord reads from reader until it hits a delimiter, such as whitespace or
// a parenthesis.
func getWord(reader *bufio.Reader) string {
	result := make([]rune, 0)

loop:
	for {
		next, err := peekRune(reader)
		if err != nil {
			break
		}

		switch next {
		case '(':
			fallthrough
		case ')':
			fallthrough
		case ' ':
			fallthrough
		case '\t':
			fallthrough
		case '\n':
			fallthrough
		case '\r':
			break loop
		default:
			reader.ReadRune()
			result = append(result, next)
		}
	}

	return string(result)
}

// peekRune returns the next rune in r without reading it.
func peekRune(r *bufio.Reader) (rune, error) {
	// Runes are at most 4 bytes long.
	buf, err := r.Peek(4)
	if len(buf) == 0 {
		return 0xFFFD, err
	}

	return bytes.Runes(buf)[0], nil
}
