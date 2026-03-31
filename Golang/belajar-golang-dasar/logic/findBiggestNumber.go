package main
import "fmt"

func main(){
	angka := [...] int {19, 100, 21, 31, 3, 4}
	temp := angka[0]

	for i:= 0; i< len(angka); i++{
		if(angka[i] > temp){
			temp = angka[i]
		}
	}
	fmt.Print(temp)
}