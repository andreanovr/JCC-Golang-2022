package main

import (
	"fmt"
	"strings"
)

func main() {
	// Soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println("\nSoal 1")
	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// Soal 2
	introduce := func(nama string, jenisKelamin string, pekerjaan string, umur string) string {
		var panggilan = ""
		if jenisKelamin == "laki-laki" {
			panggilan = "Pak"
		} else {
			panggilan = "Bu"
		}
		return panggilan + " " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}

	fmt.Println("\nSoal 2")
	var john = introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)
	var sarah = introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	// Soal 3
	var buahFavorit = func(nama string, buahBuahan ...string) string {
		kalimat := "halo nama saya " + strings.ToLower(nama) + " dan buah favorit saya adalah "
		for _, buah := range buahBuahan {
			kalimat += "\"" + buah + "\", "
		}
		return strings.Trim(kalimat, ", ")
	}
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	buahFavoritJohn := buahFavorit("John", buah...)
	fmt.Println("\nSoal 3")
	fmt.Println(buahFavoritJohn)

	// Soal 4
	fmt.Println("\nSoal 4")
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	var tambahDataFilm = func(str ...string) {
		var film = map[string]string{
			"judul":      str[0],
			"durasi":     strings.Trim(str[1], " jam"),
			"genre":      str[2],
			"tahunRilis": str[3],
		}
		dataFilm = append(dataFilm, film)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return (panjang + lebar) * 2
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}
