package main

import (
	"log"
	"net/http"

	"belajar-golang-api-basic/handlers"
)

func main() {

	http.HandleFunc("/users", handlers.CreateUser)
	http.HandleFunc("/getUsers", handlers.CallUser)

	log.Println("Route /users berhasil didaftarkan")
	log.Println("Route /getUsers berhasil didaftarkan")
	log.Println("Server berjalan di http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}