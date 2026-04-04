package main
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)


	var days int64 = 5
	var travel int64 = 2     

	scanner.Scan()
	numStr := scanner.Text()
	num, _ := strconv.ParseInt(numStr, 10, 64)

    for i:=int64(0); i<num; i++ {
        scanner.Scan()
        inputStr := scanner.Text()
        input, _ := strconv.ParseInt(inputStr, 10, 64) 
		
		results := input*days*travel
		fmt.Println(results)
	}
}