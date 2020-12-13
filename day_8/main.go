package main

import (
	"fmt"
	"strings"

	"github.com/aaronjprice/advent_of_code/lib"
)

func main() {
	instructions := parseInput("input.txt")

	accumulator, _ := run(instructions)
	fmt.Println(accumulator)

	fmt.Println(findAccAfterTerm(instructions))
}

type instruction interface {
	execute(state) state
}

type acc int

func (a acc) execute(s state) state {
	s.accumulator += int(a)
	s.nextInstruction++
	return s
}

type jmp int

func (j jmp) execute(s state) state {
	s.nextInstruction += int(j)
	return s
}

type nop int

func (n nop) execute(s state) state {
	s.nextInstruction++
	return s
}

func parseInput(filename string) []instruction {
	lines := lib.ReadLines(filename)
	instructions := []instruction{}
	for _, l := range lines {
		instructions = append(instructions, parseLine(l))
	}
	return instructions
}

func parseLine(s string) instruction {
	parts := strings.Split(s, " ")
	switch parts[0] {
	case "acc":
		return acc(lib.ToInt(parts[1]))
	case "jmp":
		return jmp(lib.ToInt(parts[1]))
	case "nop":
		return nop(lib.ToInt(parts[1]))
	default:
		panic("unrecognised command: " + parts[0])
	}
}

type state struct {
	accumulator     int
	nextInstruction int
}

func run(instructions []instruction) (int, bool) {
	s := state{}
	instructionsHistory := []int{}
	for {
		instructionsHistory = append(instructionsHistory, s.nextInstruction)
		s = instructions[s.nextInstruction].execute(s)
		if contains(instructionsHistory, s.nextInstruction) {
			return s.accumulator, false
		} else if s.nextInstruction >= len(instructions) {
			return s.accumulator, true
		}
	}
}

func findAccAfterTerm(instructions []instruction) int {
	for i := 0; i < len(instructions); i++ {
		switch t := instructions[i].(type) {
		case jmp:
			instructions[i] = nop(t)
			if accumulator, terminates := run(instructions); terminates {
				return accumulator
			}
			instructions[i] = t
		case nop:
			instructions[i] = jmp(t)
			if accumulator, terminates := run(instructions); terminates {
				return accumulator
			}
			instructions[i] = t
		default:
		}
	}
	panic("Couldn't find terminating program")
}

func contains(slice []int, target int) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
