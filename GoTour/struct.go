package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	// fmt.Println(Vertex{1, 2})

	// v := Vertex{1, 2}
	// p := &v
	// p.x = 1e9
	// fmt.Println(v)

	fmt.Println(v1, p, v2, v3)
}
