package main

import (
	"fmt"
)

type Rect struct {
	Width, Height int
}

func (r Rect) Area() int {
	return r.Width * r.Height
}

func (r Rect) String() string {
	return fmt.Sprintf("Rect ==> width=%v, height=%v", r.Width, r.Height)
}

func main() {
	r := Rect{22, 89}
	fmt.Printf("Area: %v\n", r.Area())
	fmt.Println(r)
}
