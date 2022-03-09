package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Soal 1
	var panjangPersegiPanjang = "8"
	var lebarPersegiPanjang = "5"

	var alasSegitiga = "6"
	var tinggiSegitiga = "7"

	// Converting
	panjang, _ := strconv.ParseInt(panjangPersegiPanjang, 10, 64)
	lebar, _ := strconv.ParseInt(lebarPersegiPanjang, 10, 64)
	alas, _ := strconv.ParseInt(alasSegitiga, 10, 64)
	tinggi, _ := strconv.ParseInt(tinggiSegitiga, 10, 64)

	var kelilingPersegiPanjang int64 = (panjang + lebar) * 2
	var luasSegitiga int64 = alas * tinggi / 2

	fmt.Println("\nSoal 1")
	fmt.Println("Keliling Persegi Panjang : ", kelilingPersegiPanjang, "\nLuas Segitiga : ", luasSegitiga)

	// Soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("\nSoal 2 \nNilai John A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("\nSoal 2 \nNilai John B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("\nSoal 2 \nNilai John C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("\nSoal 2 \nNilai John D")
	} else {
		fmt.Println("\nSoal 2 \nNilai John E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("Nilai Doe A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Nilai Doe B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Nilai Doe C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Nilai Doe D")
	} else {
		fmt.Println("Nilai Doe E")
	}

	// Soal 3
	var tanggal = 29
	var bulan = 11
	var tahun = 1999

	var bulanConvert = ""
	switch bulan {
	case 1:
		bulanConvert = "Januari"
	case 2:
		bulanConvert = "Februari"
	case 3:
		bulanConvert = "Maret"
	case 4:
		bulanConvert = "April"
	case 5:
		bulanConvert = "Mei"
	case 6:
		bulanConvert = "Juni"
	case 7:
		bulanConvert = "Juli"
	case 8:
		bulanConvert = "Agustus"
	case 9:
		bulanConvert = "September"
	case 10:
		bulanConvert = "Oktober"
	case 11:
		bulanConvert = "November"
	case 12:
		bulanConvert = "Desember"
	}

	dateStr := strconv.Itoa(tanggal) + " " + bulanConvert + " " + strconv.Itoa(tahun)
	fmt.Println("\nSoal 3")
	fmt.Println(dateStr)

	// Soal 4
	var gen = ""
	switch {
	case tahun >= 1995 && tahun <= 2015:
		gen = "Saya adalah Generasi Z\n"
	case tahun >= 1980 && tahun <= 1994:
		gen = "Saya adalah Generasi Y (Millenials)\n"
	case tahun >= 1965 && tahun <= 1979:
		gen = "Saya adalah Generasi X\n"
	case tahun >= 1944 && tahun <= 1964:
		gen = "Saya adalah Baby Boomer\n"
	}
	fmt.Println("\nSoal 4")
	fmt.Println(gen)
}
