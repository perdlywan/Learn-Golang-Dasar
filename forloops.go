package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	names := []string{"perdly", "Setiawan", "aja"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range collection
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	// don't need index
	for _, name := range names {
		fmt.Println(name)
	}
}
