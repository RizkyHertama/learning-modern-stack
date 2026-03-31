package main
import (
	"bufio" //package implements buffered I/O.
	"fmt" 
	"os" //provides a platform-independent interface to operating system functions
	"strconv" //package for string conversions
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input angka : ")	
	scanner.Scan()
	angkaStr := scanner.Text()
	angka, err := strconv.ParseInt(angkaStr, 10,64)

	if err != nil {
		fmt.Println("Input bukan angka yang valid!")
		return
	}

	if angka % 2 == 0{
		fmt.Println("Genap")
	}else{
		fmt.Println("Ganjil")
	}


}