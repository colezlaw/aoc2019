package main

import (
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

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			program := parse(string(p))
			program[1] = noun
			program[2] = verb

			result := computer(program)
			if result[0] == 19690720 {
				log.Printf("noun %d, verb %d, answer %d", noun, verb, noun*100+verb)
			}
		}
	}
}
