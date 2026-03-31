package main
import "fmt"

// func main(){
// 	angka := [...] int {19, 100, 21, 31, 3, 4, 21, 4}

// 	for i:=0; i<len(angka); i++{
// 		for j:=0; j<i; j++{
// 			if(angka[i] == angka[j]){
// 				fmt.Println("Duplicate number ", angka[i])	
// 			}
// 		}
// 	}
// } 


func findDuplicate(nums []int) []int{

	// Hashmap
	cariDuplikat := make(map[int]bool)
	alreadyReported := make(map[int]bool)
	var results []int

	for i:=0; i < len(nums); i++ {
		temp := nums[i]
		
		if cariDuplikat[temp]{
			if !alreadyReported[temp] {
                results = append(results, temp)
                alreadyReported[temp] = true
            }
		} else{
			cariDuplikat[temp]  = true
		}

	}
	
	// Return -1 (or any indicator) if no duplicates are found
	return results

}


func main(){
	angka := [] int {19, 100, 21, 31, 3, 4, 21, 4}
	duplikat := findDuplicate(angka)

	if len(duplikat) > 0{
		for i:= 0; i<len(duplikat); i++{
			fmt.Println("Duplicate found : ", duplikat[i])
		}
	}else{
		fmt.Println("No duplicate found.")
	}
}

