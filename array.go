package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Per"
	names[1] = "Dly"

	fmt.Println(names[0])

	var values = [...]int{
		10,
		15,
		20,
	}

	values[2] = 50
	fmt.Println(values[2])
}
