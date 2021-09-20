package main

import "fmt"

type Vector struct {
	X, Y int
}

func main() {
	// var m map[string]bool = make(map[string]bool)
	// var m2 = make(map[int]bool)
	m3 := make(map[int]bool)
	fmt.Println(m3)

	m4 := make(map[Vector]string)
	fmt.Println(m4)
}
