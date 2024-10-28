package main

import "fmt"

func main() {
	// names := []string{"Perdly", "Setiawan", "ok", "tes", "Deck", "aja"}

	// slice1 := names[4:6]
	// fmt.Println(slice1)

	// slice2 := names[:3]
	// fmt.Println(slice2)

	// slice3 := names[3:]
	// fmt.Println(slice3)

	// slice4 := names[:]
	// fmt.Println(slice4)

	days := [...]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	// daysSlice1 := days[5:]
	// daysSlice1[0] = "Jumat"
	// daysSlice1[1] = "Sabtu"
	// fmt.Println(days)

	// daySlice2 := append(daysSlice1, "baru")
	// daySlice2[0] = "halo"
	// fmt.Println(daySlice2)
	// fmt.Println(days)

	var newSlice = make([]string, 2, 5)
	newSlice[0] = "Per"
	newSlice[1] = "Dly"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "tes")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println("from", fromSlice)
	fmt.Println("to", toSlice)
}
