package main

import "fmt"

func main() {
	name := "perdly"

	switch name {
	case "perdly":
		fmt.Println("halo perdly")
	case "setiawan":
		fmt.Println("halo setiawan")
	default:
		fmt.Println("halo boleh kenalan tidak")
	}

	// switch with short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	// switch tanpa kondisi
	name = "asi"
	switch {
	case len(name) > 10:
		fmt.Println("Nama Terlalu Panjang")
	case len(name) > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
