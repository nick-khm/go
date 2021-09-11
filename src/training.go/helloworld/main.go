package main

import (
	"fmt"

	"helloworld.com/greetings/src/training.go/helloworld/input/player"
)

func main() {
	var p1 player.Player
	p1.Name = "Bob"
	p1.Age = 10

	fmt.Printf("Player 1: %v\n", p1)
	fmt.Printf("p1.Name=%v, p1.Age=%v\n", p1.Name, p1.Age)

	a := player.Avatar{"http://lol.com/ag.jpg"}
	fmt.Printf("avatar: %v\n", a)

	p2 := player.Player{
		Name: "John",
		Age:  25,
		Avatar: player.Avatar{
			Url: "http://domain/a.jpg",
		},
	}
	fmt.Printf("Player 2: %v\n", p2)

	p3 := player.New("Bobette")
	p3.Avatar = a
	fmt.Printf("Player 3: %v\n", p3)
	a.Url = "totourl"
	fmt.Printf("Player 3: %v\n", p3)
}
