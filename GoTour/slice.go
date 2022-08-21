package main

import "fmt"

func slicePointer() {

	names := [4]string{
		"john", "paul", "george", "ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "xxx"

	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// length  => slice 的长度
// cap  => slice 的第一个 到 原数组结束的长度
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sliceLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func createSliceWithMake() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5) // type length capacity
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func appendToSlice() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 5, 6)
	s1 := []int{0, 1, 2, 3, 4}
	printSlice(s)
	printSlice(s1)
}

func rangeSlice() {
	var pow = []int{1, 2, 4, 8, 16, 32, 62, 128}

	for i, v := range pow {
		fmt.Println("2**%d = %d \n", i, v)
	}
}

func Pic(dx, dy int) [][]uint8 {

	r := make([][]uint8, dy)
	for i := range r {
		r[i] = make([]uint8, dx)

		for j := range r[i] {
			r[i][j] = uint8(10 ^ j - i)
		}
	}
	return r
}

func main() {

	primes := [6]int{2, 3, 4, 7, 11, 13}

	var s []int = primes[1:4]

	fmt.Println(s)

	slicePointer()
	sliceLiterals()

	sliceLengthAndCapacity()

	createSliceWithMake()

	appendToSlice()

	rangeSlice()

}
