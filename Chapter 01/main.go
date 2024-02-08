package main

import "fmt"

type Point struct {
	X int
	Y int
}

func hello() {
	fmt.Println("Hello World")
}

func inc1(i int) int {
	return i + 1
}

func inc2(i int) (res int) {
	res = i + 1
	return
}

func sum(i, j int) int {
	return i + j
}

func calc(i int) (int, int) {
	return i * i, i + i
}

func inc(i *int) {
	*i = *i + 1
}

func main() {
	// Pointer example
	var pi *int

	i := 27
	pi = &i
	fmt.Println(*pi)

	*pi = 18
	fmt.Println(*pi)

	// Struct example
	p := Point{27, 5}
	fmt.Println(p)
	p.X = 18
	fmt.Println(p)

	pp := &p
	fmt.Println(*pp)

	i = 5
	pi = &i
	inc(pi)
	fmt.Println(i)
}
