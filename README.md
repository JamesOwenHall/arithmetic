# Arithmetic

Arithmetic is an interpreter for a Lisp-like language for basic arithmetic.

## Syntax

    (<operator> [arg [...])

Operations are nestable, so you can have equations like this:

    (+ 1 (+ 2 -4 6))

## Operators

|Symbol|   Meaning    |
|:----:|:------------:|
|  +   |   Addition   |
|  -   | Subtraction  |
|  *   |Multiplication|
|  /   |   Division   |

## Usage

    ./arithmetic "(+ 1 (+ 2 -4 6))"

Or you can use the "-i" flag to enter the REPL.

    ./arithmetic -i
    > (+ 99.332 -1)
    98.332
    > exit

## Examples

### Approximating e

    ./arithmetic "(+ 1 (/ 1 1) (/ 1 2) (/ 1 (* 2 3)) (/ 1 (* 2 3 4)) (/ 1 (* 2 3 4 5)))"
