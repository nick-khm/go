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

func processPerson(i interface{}) {
	switch p := i.(type) {
	case Person:
		fmt.Printf("We have a person=%v\n", p)
	case Chef:
		fmt.Printf("We have a chef=%v. Let him cook!", p)
		p.Cook()
	default:
		fmt.Printf("Unknown type=%T, value=%v\n", p, p)
	}
}

func main() {
	var x interface{} = Person{"Bob", 10}
	fmt.Printf("x type=%T, data=%v\n", x, x)

	s, ok := x.(string)
	fmt.Printf("Person as string ok? %v. value=%v\n", ok, s)
	n, ok := x.(int)
	fmt.Printf("Person as int ok? %v. value=%v\n", ok, n)

	processPerson(x)

	x = Chef{
		Stars: 4,
		Person: Person{
			Name: "Alice",
			Age:  22,
		},
	}

	processPerson(x)

	x = 3
	processPerson(x)

	x = "John"
	processPerson(x)

}
