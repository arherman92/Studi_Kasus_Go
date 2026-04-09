package main

import "fmt"

func main() {
	var nilai1, nilai2, nilai3 float64

	fmt.Println("Masukkan nilai 1 : ")
	fmt.Scanln(&nilai1)

	fmt.Println("Masukkan nilai 2 : ")
	fmt.Scanln(&nilai2)

	fmt.Println("Masukkan nilai 3 : ")
	fmt.Scanln(&nilai3)

	rataRata := (nilai1 + nilai2 + nilai3) / 3

	fmt.Println("Rata - rata :", rataRata)

	if rataRata >= 75 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tida Lulus")
	}

}
