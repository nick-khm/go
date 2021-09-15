package main

import (
	"fmt"
	"os"

	"helloworld.com/greetings/src/training.go/hangman/hangman"
)

func main() {
	g := hangman.New(8, "Golang")

	hangman.DrawWelcome()
	guess := ""
	for {
		hangman.Draw(g, guess)
		switch g.State {
		case "win", "lost":
			os.Exit(0)
		}
		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Cound not read from terminal: %v", err)
			os.Exit(1)
		}
		guess = l
	}

}
