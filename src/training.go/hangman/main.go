package main

import (
	"fmt"
	"os"

	"helloworld.com/greetings/src/training.go/hangman/hangman"
	"helloworld.com/greetings/src/training.go/hangman/hangman/dictionary"
)

func main() {
	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load the dictionary: %v\n", err)
		os.Exit(1)
	}

	g := hangman.New(8, dictionary.PickWord())

	hangman.DrawWelcome()
	guess := ""
	for {
		hangman.Draw(g, guess)
		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}
		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Cound not read from terminal: %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}

}
