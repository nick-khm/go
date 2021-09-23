package main

import "fmt"

type Vector struct {
	X, Y int
}

func main() {
	m := map[string]int{
		"Bob":     10,
		"Alice":   15,
		"Bobette": 20,
		"John":    7,
	}

	for name, age := range m {
		fmt.Printf("name=%v, age=%v\n", name, age)
	}

	i := 1
	for name := range m {
		fmt.Printf("name=%v\n", name)
		m[name] = i
		i++
	}
	fmt.Println(m)

}
