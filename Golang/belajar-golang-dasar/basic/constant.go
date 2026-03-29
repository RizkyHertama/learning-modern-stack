package main
import "fmt"

func main(){
	// Const merupakan variabel yg nilainya tetap / statis

	const firstName string = "Rizky"
	const lastName = "Hertama"
	
	//error
	// firstName = "budi"
	// lastName = "koko" 

	const (
		animalName = "Kucing"
		animalSound = "Meow"
	)

	fmt.Println(animalName)
	fmt.Println(animalSound)
}