package main

import "fmt"

func main(){
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	name := []string {"Eko", "Kurniawan", "Khanedy"}
	for i := 0; i < len(name); i++{
		fmt.Println(name[i])
	}

	for index, name := range name{
		fmt.Println("Index", index, "=", name)
	}
}