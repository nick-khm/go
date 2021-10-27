package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	color string
}

func (d Dog) Speak() string {
	return "woof!"
}

type Cat struct {
	jumpHeight int
}

func (c Cat) Speak() string {
	return "miao!"
}

func main() {
	animals := []Animal{
		Dog{"white"},
		Cat{2},
	}

	for _, a := range animals {
		fmt.Println(a.Speak())
		fmt.Printf("Type of animal? %T\n", a)

		// type assertion
		// t, ok := a.(Dog)
		// fmt.Printf("Type assertion value=%v, ok=%v\n", t, ok)

		if t, ok := a.(Dog); ok {
			fmt.Printf("We have a dog! Color=%v\n", t.color)
		} else {
			fmt.Println("It's not a dog here...")
		}
	}

	fmt.Println("------------------------")

	for _, a := range animals {
		describeAnimal(a)
	}
}

func describeAnimal(a Animal) {
	switch v := a.(type) {
	case Dog:
		fmt.Printf("We have a dog. Color=%v\n", v.color)
	case Cat:
		fmt.Printf("We have a cat. Jump height=%v\n", v.jumpHeight)
	}
}
