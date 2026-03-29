package main
import "fmt"

func main(){
	names := [...] string{"Rizky", "Hertama", "Joko", "Budi", "Nugroho", "Widodo"}

	slice1:= names[4:6]
	fmt.Println(slice1)

	slice2:= names[:3]
	fmt.Println(slice2)

	slice3:= names[3:]
	fmt.Println(slice3)

	// slice4:= names[:]
	var slice4 []string = names[:]
	fmt.Println(slice4)

	days := [...] string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	
	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2,5) //Panjang 2 dan kapasitas 5
	newSlice[0] = "Eko"
	newSlice[1] = "Eko"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Eko")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	intArray := [...]int{1, 2, 3, 4, 5}
	intSlice := []int{1, 2, 3, 4, 5}
	
	fmt.Println(intArray)
	fmt.Println(intSlice)
}