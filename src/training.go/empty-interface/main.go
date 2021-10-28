package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Cooker interface {
	Cook()
}

type Chef struct {
	Person
	Stars int
}

func (c Chef) Cook() {
	fmt.Printf("I'm cooking with %v stars\n", c.Stars)
}

func main() {
	var x interface{} = Person{"Bob", 10}
	fmt.Printf("x type=%T, data=%v\n", x, x)
}
