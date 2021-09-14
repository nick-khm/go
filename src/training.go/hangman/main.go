package main

import (
	"fmt"

	"helloworld.com/greetings/src/training.go/hangman/hangman"
)

func main() {
	g := hangman.New(8, "Golang")
	fmt.Println(g)
}
