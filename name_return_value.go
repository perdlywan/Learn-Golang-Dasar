package main

import "fmt"

func returnValue()(firstName, lastName string, age int) {
	firstName = "perdly"
	lastName = "setiawan"
	age = 17

	return firstName, lastName, age
}

func main() {
	a, b, c := returnValue();
	fmt.Println(a, b, c)
}