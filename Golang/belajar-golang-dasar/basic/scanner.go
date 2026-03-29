package main
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	//1. Input name (String)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	//2. Input salary (Float)
	fmt.Print("Enter your salary: ")
	scanner.Scan()
	salaryStr := scanner.Text()
	salary, _ := strconv.ParseFloat(salaryStr, 64)
	
	//3. Input age (Int)
	fmt.Print("Enter your age: ")
	scanner.Scan()
	ageStr := scanner.Text()
	age, _ := strconv.ParseInt(ageStr, 64)

	//Output
	fmt.Println("--- Employee Details ---")
	fmt.Println("Name ", name)
	fmt.Println("Age", age)
	fmt.Println("Salary", salary)
}