package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	log := `LOG [2026-05-20] User: @budi_01 Berhasil melakukan transaksi TX99281 senilai Rp. 550.000 via Bank.`

	//1. Ekstrak ID Transaksi
	reID := regexp.MustCompile(`TX\d{5}`)
	idMatch := reID.FindString(log)

	// 2. Ekstrak Nominal (ambil angka setelah Rp. dan hapus titiknya)
	reDuit := regexp.MustCompile(`Rp\.\s?([\d\.]+)`)

	duitMatch := reDuit.FindStringSubmatch(log)

	bersih := ""

	if len(duitMatch) > 1 {
		bersih = strings.ReplaceAll(duitMatch[1], ".", "")
	}

	fmt.Println("ID transaksi:", idMatch)
	fmt.Println("Nominal:", bersih)
}
