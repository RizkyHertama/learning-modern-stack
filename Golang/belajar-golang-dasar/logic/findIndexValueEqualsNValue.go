/*
Given array {3,4,1,8,9,5}
N = 7 
Which index value total = N?
Answer : index(0, 1)
*/

package main
import "fmt"

func main(){
	n := 7
	angka := [...]int{3,4,1,8,9,5}

	for i:=0; i< len(angka); i++{
		for j:=0; j<i; j++{
			totalIdx := angka[i] + angka[j]
			if(totalIdx == n){
				fmt.Print("Index yang menghasilkan nilai dari N = ", n , " adalah " , i, " dan " ,j)
			}	
		}
	}
}
