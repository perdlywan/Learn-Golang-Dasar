package main

import "fmt"

func main() {
	type NoKTP string

	var ktpSaya NoKTP = "111111"
	fmt.Println(ktpSaya)
	fmt.Println(NoKTP("222222"))
}
