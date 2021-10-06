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

	d.Add("golang", "wonderful language")
	entry, _ := d.Get("golang")
	fmt.Println(entry)
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
		os.Exit(1)
	}
}
