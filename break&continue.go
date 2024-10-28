package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println(i)
	}

	// mencari bilangan ganjil
	for i := 0; i <= 50; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}
}
