package main

import (
	"fmt"
)

func getUserInput() (s string, score float64) {
	fmt.Printf("qing shu ru yao shou suo de guan jian zi:")
	_, err := fmt.Scanln(&s)
	if err != nil {
		s = "default"
	}
	fmt.Printf("qing shu ru fen shu:")
	_, err = fmt.Scanln(&score)
	if err != nil {
		score = 0
	}
	return
}

func main() {
	title, score := getUserInput()
	fmt.Printf("%v %v", title, score)
}
