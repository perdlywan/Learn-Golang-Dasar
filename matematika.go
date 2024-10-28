package main

import "fmt"

func main() {
	// Operasi Matematika Dasar
	var a = 10
	var b = 15

	var c = a + b
	fmt.Println(c)

	// assignment operator
	var i = 9
	i += 10
	fmt.Println(i)

	// unary operator
	var g = 2
	g++
	fmt.Println(g)
	g++
	fmt.Println(g)
	g--
	fmt.Println(g)
}
