package main;
import "fmt";

func main(){
	
	//1. Full declaration variable (with data types)
	var fruits string = "Apple"
	var score int = 100
	var salary float64 = 3.53

	fmt.Println(fruits, score, salary)

	//2. Short
	name := "Rizky Hertama"
	fmt.Println(name);

	name = "Jono"
	fmt.Println(name);

	//3. Multiple declare variable
	var (
		firstName = "Revaldy"
		middleName = "Kina"
		lastName = "Hertama"
	)

	fmt.Println(firstName, middleName, lastName);

}