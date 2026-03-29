package main
import "fmt"


func main(){
	type NoKTP string

	var ktpIky NoKTP = "12231213122"

	var contoh string = "2222222222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpIky)
	fmt.Println(contohKTP)
}