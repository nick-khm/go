package main

import "fmt"

type Instrumenter interface {
	Play()
}

type Amplifier interface {
	Connect(amp string)
}

type Guitar struct {
	Strings int
}

func (g Guitar) Play() {
	fmt.Printf("Tzouiiing with %d strings\n", g.Strings)
}

func (g Guitar) Connect(amp string) {
	fmt.Printf("Connected to %v\n", amp)
}

type Piano struct {
	Keys int
}

func (p Piano) Play() {
	fmt.Printf("Plip plip %d keys\n", p.Keys)
}

func main() {
	var instr Instrumenter
	instr = &Guitar{5}
	instr.Play()

	instr = &Piano{88}
	instr.Play()

	guitar := Guitar{12}
	guitar.Connect("Marshall")
	guitar.Play()
}
