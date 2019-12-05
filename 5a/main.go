package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func input() int {
	return 1
}

func output(v int) {
	fmt.Println(v)
}

func computer(program []int) []int {
	result := make([]int, len(program))
	copy(result, program)

	pc := 0
	for {
		opcode := result[pc] % 100
		modes := fmt.Sprintf("%03d", result[pc]/100)

		switch opcode {
		case 99:
			return result
		case 1:
			addend1 := result[pc+1]
			if modes[2] == '0' {
				addend1 = result[addend1]
			}
			addend2 := result[pc+2]
			if modes[1] == '0' {
				addend2 = result[addend2]
			}
			result[result[pc+3]] = addend1 + addend2
			pc += 4
		case 2:
			factor1 := result[pc+1]
			if modes[2] == '0' {
				factor1 = result[factor1]
			}
			factor2 := result[pc+2]
			if modes[1] == '0' {
				factor2 = result[factor2]
			}
			result[result[pc+3]] = factor1 * factor2
			pc += 4
		case 3:
			result[result[pc+1]] = input()
			pc += 2
		case 4:
			target := pc + 1
			if modes[2] == '0' {
				target = result[pc+1]
			}
			output(result[target])
			pc += 2
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

	result := computer(program)
	log.Printf("%d\n", result)
}
