package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Location constants
const (
	LoxEmpty int = iota
	LocOrigin
	LocWire1
	LocWire2
	LocBoth
)

func parse(in string) []string {
	return strings.Split(in, ",")
}

func intersect(wire1, wire2 []string) int {
	field := map[int]map[int]int{
		0: map[int]int{0: LocOrigin},
	}

	x, y := 0, 0

	for _, move := range wire1 {
		dir := move[0]
		distance, _ := strconv.Atoi(move[1:])

		for i := 0; i < distance; i++ {
			switch dir {
			case 'U':
				y++
			case 'D':
				y--
			case 'L':
				x--
			case 'R':
				x++
			}
			if _, ok := field[x]; !ok {
				field[x] = make(map[int]int)
			}
			field[x][y] = LocWire1
		}
	}

	x, y = 0, 0
	min := 1<<32 - 1

	for _, move := range wire2 {
		dir := move[0]
		distance, _ := strconv.Atoi(move[1:])

		for i := 0; i < distance; i++ {
			switch dir {
			case 'U':
				y++
			case 'D':
				y--
			case 'L':
				x--
			case 'R':
				x++
			}
			if _, ok := field[x]; !ok {
				field[x] = make(map[int]int)
			}
			if field[x][y] == LocWire1 {
				field[x][y] = LocBoth
				md := math.Abs(float64(x)) + math.Abs(float64(y))
				if int(md) < min {
					min = int(md)
				}
			}
		}
	}

	return min
}

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Open: %v", err)
	}
	defer r.Close()

	s := bufio.NewScanner(r)
	s.Scan()
	wire1 := s.Text()
	s.Scan()
	wire2 := s.Text()

	fmt.Printf("Min distance %d\n", intersect(parse(wire1), parse(wire2)))
}
