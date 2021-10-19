package main

import (
	"fmt"
	"os"

	"helloworld.com/greetings/src/training.go/dictionary/dictionary"
)

func main() {
	d, err := dictionary.New("./badger")
	handleErr(err)
	defer d.Close()

	/*d.Add("java", "dinosour like language")
	words, entries, _ := d.List()
	for _, word := range words {
		fmt.Println(entries[word])
	}*/

	d.Remove("Java")
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
		os.Exit(1)
	}
}
