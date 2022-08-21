package main

import "fmt"

type Vertex struct {
	Lat, Lng float64
}

var m map[string]Vertex

func mapLiterals() {

	var m = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

}

func MutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("the value", m["Answer"])

	m["Answer"] = 48
	fmt.Println("the value", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v1, ok1 := m["Answer"]
	fmt.Println("The value:", v1, "Present?", ok1)

}

func main() {

	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	mapLiterals()

	MutatingMaps()
}
