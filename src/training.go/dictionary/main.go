package main

import (
	"flag"
	"fmt"
	"os"

	"helloworld.com/greetings/src/training.go/dictionary/dictionary"
)

func main() {
	action := flag.String("action", "list", "Action to perform on the dictionary")

	d, err := dictionary.New("./badger")
	handleErr(err)
	defer d.Close()

	flag.Parse()

	switch *action {
	case "list":
		actionList(d)

	default:
		fmt.Printf("Unknown action %v\n", *action)
	}
}

func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	handleErr(err)
	fmt.Println("Dictionary content")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
		os.Exit(1)
	}
}
