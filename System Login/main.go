package main

import (
	"fmt"
)

func main() {
	// Data user (hardcode)

	var correctUsername = "admin"
	var correctPassword = "12345"

	// Input User
	var userName, password string

	fmt.Print("Masukkan username: ")
	fmt.Scanln(&userName)

	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&password)

	// validasi login

	if userName == correctUsername && password == correctPassword {
		fmt.Println("Login berhasil")
	} else {
		fmt.Println("Login Gagal")
	}
}
