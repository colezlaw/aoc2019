package main

import (
	"fmt"
	"log"
)

func isValidPassword(pw string) bool {
	if len(pw) != 6 {
		log.Printf("not len 6 %s", pw)
		return false
	}

	for i := 0; i < len(pw); i++ {
		if pw[i] < '0' || pw[i] > '9' {
			log.Printf("not digits %s", pw)
			return false
		}
	}

	prior := -1
	repeat := false

	for i := 0; i < len(pw); i++ {
		if int(pw[i])-48 == prior {
			log.Printf("repeat true")
			repeat = true
		}
		if int(pw[i])-48 < prior {
			log.Printf("%d < %d (%s)", int(pw[i]-48), prior, pw)
			repeat = false
		}
		prior = int(pw[i]) - 48
	}

	return repeat
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
