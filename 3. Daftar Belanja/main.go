package main

import "fmt"

func main() {
	// Tujuan
	// Menyimpan Data list
	// Menampilkan data mengunakan for

	// slice (array dinamis) []string
	daftarBelanja := []string{"Beras", "Minyak", "Telur"}

	//Tambah item/data disini dengan menggunakan append() dari variavel Slice
	daftarBelanja = append(daftarBelanja, "Gula", "Sayur Labu")

	//Tampilkan semua menggunakan for dan range untuk looping
	for i, item := range daftarBelanja {
		fmt.Println(i+1, ".", item)
	}
}
