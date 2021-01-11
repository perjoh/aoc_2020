package main

import (
	"aoc2020/util"
	"fmt"
)

type valueParser interface {
	parseValue(e *evaluator) int
}

type valueParserDefault struct {
}

func (valueParserDefault) parseValue(e *evaluator) int {
	if e.char() == '(' {
		e.step()
		return e.parseExpr()
	}
	value := int(e.char() - '0')
	e.step()
	return value
}

type valueParserFunky struct {
}

func (valueParserFunky) parseValue(e *evaluator) int {
	value := valueParserDefault{}.parseValue(e)
	if !e.eof() {
		if e.char() == '+' {
			e.step()
			value = value + e.parseValue()
		}
	}
	return value
}

type evaluator struct {
	expr string
	pos  int
	vp   valueParser
}

func (e evaluator) char() byte {
	return e.expr[e.pos]
}

func (e *evaluator) step() {
	e.pos++
	if !e.eof() && e.char() == ' ' {
		e.step()
	}
}

func (e evaluator) eof() bool {
	return !(e.pos < len(e.expr))
}

func (e *evaluator) parseValue() int {
	return e.vp.parseValue(e)
}

func (e *evaluator) parseExpr() int {
	value := 0
	for !e.eof() {
		c := e.char()
		switch c {
		case ')':
			e.step()
			return value
		case '+':
			e.step()
			value = value + e.parseValue()
		case '*':
			e.step()
			value = value * e.parseValue()
		default:
			value = e.parseValue()
		}
	}
	return value
}

func eval(expr string) int {
	e := evaluator{expr: expr, vp: &valueParserDefault{}}
	return e.parseExpr()
}

func part1(input []string) int {
	sum := 0
	for _, expr := range input {
		sum += eval(expr)
	}
	return sum
}

func evalFunky(expr string) int {
	e := evaluator{expr: expr, vp: &valueParserFunky{}}
	return e.parseExpr()
}

func part2(input []string) int {
	sum := 0
	for _, expr := range input {
		sum += evalFunky(expr)
	}
	return sum
}

func main() {
	input := util.ReadInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
