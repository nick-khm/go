package main

import (
	"fmt"
	"os"

	"helloworld.com/greetings/src/training.go/hangman/hangman"
)

func main() {
	g := hangman.New(8, "Golang")
	fmt.Println(g)

	l, err := hangman.ReadGuess()
	if err != nil {
		fmt.Printf("Cound not read from terminal: %v", err)
		os.Exit(1)
	}
	fmt.Println(l)
}
