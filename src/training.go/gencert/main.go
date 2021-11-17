package main

import (
	"fmt"
	"os"

	"helloworld.com/greetings/src/training.go/gencert/cert"
	"helloworld.com/greetings/src/training.go/gencert/pdf"
)

func main() {
	c, err := cert.New("Golang programming", "Bob Dylan", "2021-11-17")
	if err != nil {
		fmt.Printf("Error during a certificate creation: %v\n", err)
		os.Exit(1)
	}

	var saver cert.Saver
	saver, err = pdf.New("output")
	if err != nil {
		fmt.Printf("Error during a certificate creation: %v\n", err)
		os.Exit(1)
	}

	saver.Save(*c)
}
