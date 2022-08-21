package main

import "fmt"

func stackDefers() {

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	stackDefers()
}
