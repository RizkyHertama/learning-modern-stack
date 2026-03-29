package main

import "fmt"

func main(){
	// name := "Eko" 
	// name := "Joko"
	// name := "Iky"
	name := "Jokowidodo"

	if name == "Eko" {
		fmt.Println("Haloo ekoo")
	}else if name == "Joko" {
		fmt.Println("Halo joko")
	}else{
		fmt.Println("Bukan eko dan joko")
	}

	length := len(name)

	if length > 5{
		fmt.Println("Nama lebih dari 5 karakter")
	}else{
		fmt.Println("Nama kurang dari 5 karakter")
	}
}