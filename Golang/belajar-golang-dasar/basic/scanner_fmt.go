package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	// var nama string
	var umur int
	var tinggi float64
	var berat float64

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input nama karyawan : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Input umur karyawan : ")
	fmt.Scanln(&umur)

	fmt.Print("Input tinggi karyawan (cm) : ")
	fmt.Scanln(&tinggi)

	fmt.Print("Input berat karyawan (kg) : ")
	fmt.Scanln(&berat)

	fmt.Println("\n====================")
	fmt.Println("Nama  :", nama)
	fmt.Println("Umur  :", umur)
	fmt.Println("Tinggi:", tinggi, "cm")
	fmt.Println("Berat :", berat, "kg")
}