package main

import "fmt"

type Vector struct {
	X, Y int
}

func main() {
	m := make(map[string]int)
	fmt.Printf("Map content %v len(%v)\n", m, len(m))

	m["hello"] = 5
	m["good bye"] = 10
	fmt.Printf("Map content %v len(%v)\n", m, len(m))
	fmt.Printf("key=hello key=%v\n", m["hello"])

	// presence of a key
	j, ok := m["toto"]
	fmt.Printf("j=%v ok=%v\n", j, ok)

	// presence of a key
	d, okk := m["hello"]
	fmt.Printf("d=%v okk=%v\n", d, okk)

	m["beatles"] = 2
	if _, ok := m["beatles"]; ok {
		fmt.Println("beatles key exists")
	}

	delete(m, "beatles")
	fmt.Printf("Map content %v len(%v)\n", m, len(m))

	m2 := m
	fmt.Printf("m2 content %v (len=%v)\n", m2, len(m2))

	m2["help"] = 44
	fmt.Printf("m content %v (len=%v)\n", m, len(m))
	fmt.Printf("m2 content %v (len=%v)\n", m2, len(m2))
}
