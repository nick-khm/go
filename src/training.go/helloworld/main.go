package main

import (
	"fmt"
)

func UpdateVal(val string) {
	val = "myValue"
}

func UpdatePointer(ptr *string) {
	*ptr = "myValue"
}

func main() {
	i := 1
	var p *int = &i

	fmt.Printf("i=%v\n", i)
	fmt.Printf("p=%v\n", p)
	fmt.Printf("*p=%v\n", *p)
	fmt.Printf("p=%v\n", p)

	fmt.Println("-----------------")

	s := "Bob"
	sPtr := &s
	s2 := *sPtr

	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPrt=%v\n", sPtr)
	fmt.Printf("*sPrt=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)

	fmt.Println("-----------------")

	s = "Alice"

	fmt.Println("Change the initial variable from Bob to Alice")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPrt=%v\n", sPtr)
	fmt.Printf("*sPrt=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)

	fmt.Println("-----------------")

	*sPtr = "Hugo"

	fmt.Println("Dereference sPrt and update it with Hugo")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPrt=%v\n", sPtr)
	fmt.Printf("*sPrt=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)

	fmt.Println("-----------------")

	UpdateVal(s)

	fmt.Println("UpdateVal => trying to change to myValue")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPrt=%v\n", sPtr)
	fmt.Printf("*sPrt=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)

	fmt.Println("-----------------")

	// UpdatePointer(&s)
	UpdatePointer(sPtr)

	fmt.Println("UpdatePointer => trying to change to myValue")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPrt=%v\n", sPtr)
	fmt.Printf("*sPrt=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)

	fmt.Println("-----------------")

}
