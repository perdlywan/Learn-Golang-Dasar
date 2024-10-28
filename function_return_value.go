package main

import "fmt"

func getMultipleValue() (string, string) {
	return "Daniel", "Joshua"
}

func getValue(name string) string {
	hello := "Hallo" + name
	return hello
}

func main() {
	fmt.Println(getValue("perd"))
	fmt.Println(getMultipleValue())
	_, secondName := getMultipleValue()
	fmt.Println(secondName)
}
