package main

import "fmt"

func main(){
		// name := "Eko" 
	// name := "Joko"
	// name := "Iky"
	name := "Jokowidodo"

	switch name{
	case "Eko" :
		fmt.Println("Haloo ekoo")
	case "Joko" :
		fmt.Println("Haloo jokoo")
	default :
		fmt.Println("Bukan eko dan joko")
	}


	length := len(name)

	switch length > 5 {
		case true : 
			fmt.Println("Nama lebih dari 5 karakter")
		case false :
			fmt.Println("Nama kurang dari 5 karakter")
	}
}