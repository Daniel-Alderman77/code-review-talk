package main

import (
	"fmt"

	"github.com/Daniel-Alderman77/code-review-talk/triangles/geometry"
)

func main() {
	a, b, c := 3, 4, 5
	fmt.Printf("A triangle with sides: %d, %d, %d", a, b, c)
	t := geometry.Triangle(a, b, c)
	fmt.Printf("The triangle is a %s", t)
}
