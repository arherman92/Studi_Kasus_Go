package main

import "fmt"

func main() {
	luas := hitungLuas(5, 3)

	fmt.Println("Luas", luas)
}

// function
func hitungLuas(panjang int, lebar int) int {
	return panjang * lebar
}
