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
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
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

func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleErr(err)
	fmt.Printf("'%v' added to the dictionary\n", word)
}

func actionDefine(d *dictionary.Dictionary, args []string) {
	entry, err := d.Get(args[0])
	handleErr(err)
	fmt.Println(entry)
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
		os.Exit(1)
	}
}
