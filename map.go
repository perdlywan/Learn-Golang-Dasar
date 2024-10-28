package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["nama"] = "perdly"
	// person["alamat"] = "kepri"

	person := map[string]string{
		"name":    "perd",
		"address": "kepri",
	}

	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Pemrograman Web 1"
	book["author"] = "perd"
	book["year"] = "2009"
	book["tes"] = "asdas"

	delete(book, "tes")

	fmt.Println(book)
}
