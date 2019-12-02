package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calcMass(mass int) int {
	return mass/3 - 2
}

func main() {
	sum := 0
	r, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Open %v", err)
	}
	defer r.Close()

	s := bufio.NewScanner(r)
	for s.Scan() {
		mass, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("Atoi: %v", err)
		}
		sum += calcMass(mass)
	}

	fmt.Printf("Total fuel required %d\n", sum)
}
