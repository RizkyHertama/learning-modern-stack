package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i:= 0; i<t; i++{
		scanner.Scan()
		line := scanner.Text()

		// Pecah baris jadi slice of strings: ["10", "12", "4"]
		parts := strings.Fields(line)

		// Inisialisasi slice kosong untuk menampung int
		var numbers []int

		for j := 0; j<len(parts); j++{
			val, _ := strconv.Atoi(parts[j])
			numbers = append(numbers, val)
		}
	
		sort.Ints(numbers)
		if len(numbers) >= 2{
			fmt.Println(numbers[len(numbers)-2])
		}
	}
}