package main

import (
	"aoc2020/util"
	"fmt"
	"strconv"
	"strings"
)

type cpu struct {
	ip        int
	acc       int
	loopGuard map[int]struct{}
	panic     bool
}

func parseInstruction(instruction string) (string, int) {
	ins := strings.Split(instruction, " ")
	val, _ := strconv.Atoi(ins[1])
	return ins[0], val
}

func (c *cpu) runInstruction(opCode string, arg int) {
	_, has := c.loopGuard[c.ip]
	if !has {
		c.loopGuard[c.ip] = struct{}{}
		switch opCode {
		case "nop":
			c.ip++
		case "acc":
			c.acc += arg
			c.ip++
		case "jmp":
			c.ip += arg
		}
	} else {
		c.panic = true
	}
}

func (c *cpu) runProgram(program []string) {
	for !c.panic && c.ip < len(program) {
		opCode, arg := parseInstruction(program[c.ip])
		c.runInstruction(opCode, arg)
	}
}

func partA(input []string) {
	c := cpu{loopGuard: make(map[int]struct{})}
	c.runProgram(input)
	fmt.Println(c.acc)
}

func patchInstruction(instruction string) string {
	if instruction[0] == 'n' {
		return strings.Replace(instruction, "nop", "jmp", 1)
	}
	return strings.Replace(instruction, "jmp", "nop", 1)
}

func partB(input []string) {
	for ip := range input {
		if input[ip][0] != 'a' {
			program := append([]string{}, input...)
			c := cpu{loopGuard: make(map[int]struct{})}
			program[ip] = patchInstruction(program[ip])
			c.runProgram(program)
			if !c.panic {
				fmt.Println(c.acc)
				break
			}
		}
	}
}

func main() {
	input := util.ReadStdin()
	partA(input)
	partB(input)
}
