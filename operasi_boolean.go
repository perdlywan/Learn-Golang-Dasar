package main

import "fmt"

func main() {
	var nilaiuts = 90
	var nilaiuas = 80

	var lulusuts = nilaiuts > 80
	var lulusuas = nilaiuas > 80

	var lulus = lulusuts && lulusuas
	fmt.Println(lulus)
}
