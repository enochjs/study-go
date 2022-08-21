package main

import (
	"fmt"
	"runtime"
	"time"
)

func timeToSaturday() {
	fmt.Print("when's saturday?\n")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("today .")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 1:
		fmt.Println("in two days")
	default:
		fmt.Println("too far away")
	}

}

func switchNoCondition() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("good morning!")
	case t.Hour() < 17:
		fmt.Println("goof afternoon.")
	default:
		fmt.Println("good evening")
	}
}

func main() {
	fmt.Print("go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("os x")
	case "linux":
		fmt.Print("linux")
	default:
		fmt.Println("%s. \n", os)
	}

	timeToSaturday()

	switchNoCondition()

}
