package main

import (
	"assignment1"
	"fmt"
	"os"
	"strconv"
)

func PrintAbsensiData(noAbsen int, absensi assignment1.Absensi) {
	fmt.Printf(`
=============DATA ABSENSI ke %d=================
Nama: %s,
Alamat: %s,
Pekerjaan: %s,
Alasan memilih kelas: %s
==========================================

`, noAbsen, absensi.Name, absensi.Address, absensi.Job, absensi.Reason)
}

func main() {
	dataAbsensi := assignment1.ProvideAbsensiDataDummy()

	if len(os.Args) < 2 {
		fmt.Println("masukkan data argument no absen")
	} else {
		argumentData := os.Args[1]
		argumentDataInt, err := strconv.Atoi(argumentData)
		if err != nil {
			panic(err)
		}

		if argumentDataInt > len(dataAbsensi) {
			fmt.Println("Data tidak ditemukan")
		} else {
			PrintAbsensiData(argumentDataInt, dataAbsensi[argumentDataInt-1])
		}
	}

}
