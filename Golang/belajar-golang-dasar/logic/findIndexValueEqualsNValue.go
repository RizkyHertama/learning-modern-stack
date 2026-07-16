/*
Given array {3,4,1,8,9,5}
N = 7 
Which index value total = N?
Answer : index(0, 1)
*/

package main
import "fmt"

func main(){
	var number int = 13;
	var values = [...] int {12, 7, 8, 3, 2, 1}
	var found bool = false;

	for i := 0; i<len(values); i++{
		for j := 0; j<i; j++{
			
			var total int = values[i] + values[j];
			if(total == number){
				fmt.Println(j, i)
				found = true;
				break;
			}else if(values[i] == number){
				fmt.Println(i)
				found = true;
				break;
			}else if(values[j] == number){
				fmt.Println(j)
				found = true;
				break;
			}
		}
	}
	if(!found){
		fmt.Println("Data From List Not Found")
	}
}


/*
Secara umum jumlah iterasi pada program tersebut itu ditentukan oleh banyaknya list values
Rumusnya n*(n-1)/2 = 6*5/2 = 15
{12, 7, 8, 3, 2, 1}

| Iterasi | i | j | values[i] | values[j] | total | Hasil 			|
|---------|---|---|-----------|-----------|-------|-----------------|
| 1       | 1 | 0 | 7         | 12        | 19    | ❌ tidak cocok |
| 2       | 2 | 0 | 8         | 12        | 20    | ❌ tidak cocok |
| 3       | 2 | 1 | 8         | 7         | 15    | ❌ tidak cocok |
| 4       | 3 | 0 | 3         | 12        | 15    | ❌ tidak cocok |
| 5       | 3 | 1 | 3         | 7         | 10    | ❌ tidak cocok |
| 6       | 3 | 2 | 3         | 8         | 11    | ❌ tidak cocok |
| 7       | 4 | 0 | 2         | 12        | 14    | ❌ tidak cocok |
| 8       | 4 | 1 | 2         | 7         | 9     | ❌ tidak cocok |
| 9       | 4 | 2 | 2         | 8         | 10    | ❌ tidak cocok |
| 10      | 4 | 3 | 2         | 3         | 5     | ❌ tidak cocok |
| 11      | 5 | 0 | 1         | 12        | 13    | ✅ cocok → cetak (0,5) |
| 12      | 5 | 1 | 1         | 7         | 8     | ❌ tidak cocok |
| 13      | 5 | 2 | 1         | 8         | 9     | ❌ tidak cocok |
| 14      | 5 | 3 | 1         | 3         | 4     | ❌ tidak cocok |
| 15      | 5 | 4 | 1         | 2         | 3     | ❌ tidak cocok |

*/
