package main

import (
	"fmt"
	"log"
)

func isValidPassword(pw string) bool {
	if len(pw) != 6 {
		return false
	}

	for i := 0; i < len(pw); i++ {
		if pw[i] < '0' || pw[i] > '9' {
			return false
		}
	}

	prior, prior2 := -1, -1
	repeat := make(map[int]bool)

	for i := 0; i < len(pw); i++ {
		if int(pw[i])-48 == prior {
			repeat[int(pw[i])-48] = true
			if int(pw[i]-48) == prior2 {
				log.Printf("long group in %s", pw)
				repeat[int(pw[i])-48] = false
			}
		}
		if int(pw[i])-48 < prior {
			return false
		}
		prior2 = prior
		prior = int(pw[i]) - 48
	}

	for _, v := range repeat {
		if v {
			return v
		}
	}

	return false
}

func main() {
	count := 0
	for i := 171309; i <= 643603; i++ {
		if isValidPassword(fmt.Sprintf("%d", i)) {
			count++
		}
	}

	fmt.Printf("Valid passwords %d\n", count)
}
