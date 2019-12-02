package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func computer(program []int) []int {
	result := make([]int, len(program))
	copy(result, program)

	pc := 0
	for {
		switch result[pc] {
		case 99:
			return result
		case 1:
			result[result[pc+3]] = result[result[pc+1]] + result[result[pc+2]]
			pc += 4
		case 2:
			result[result[pc+3]] = result[result[pc+1]] * result[result[pc+2]]
			pc += 4
		default:
			log.Fatalf("at pc %d, unknown opcode %d", pc, result[pc])
		}
	}
}

func parse(code string) []int {
	result := strings.Split(code, ",")
	program := make([]int, len(result))
	for i, x := range result {
		c, err := strconv.Atoi(x)
		if err != nil {
			log.Fatalf("unable to parse int at %d %v", i, err)
		}
		program[i] = c
	}

	return program
}

func main() {
	p, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("ReadFile: %v", err)
	}

	program := parse(string(p))
	program[1] = 12
	program[2] = 2

	result := computer(program)
	fmt.Printf("result[0] == %d\n", result[0])
}
