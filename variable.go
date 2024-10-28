package main

import "fmt"

func main() {
	var name string

	name = "perdly"
	fmt.Println(name)

	name = "perdly doang"
	fmt.Println(name)

	// deklarasi variable tanpa tipe data
	var name2 = "Budi"
	fmt.Println(name2)

	// deklarasi variable tanpa var
	name3 := "Eko"
	fmt.Println(name3)

	name3 = "andi"
	fmt.Println(name3)

	var (
		firstname  = "perdly"
		middlename = "setiawan"
		lastname   = "doang"
	)

	fmt.Println(firstname)
	fmt.Println(middlename)
	fmt.Println(lastname)
}
