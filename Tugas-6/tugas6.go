package main

import (
	"fmt"
)

func main() {
	// Soal 1
	soal1()

	// Soal 2
	soal2()

	// Soal 3
	soal3()

	// Soal 4
	soal4()
}

func soal1() {
	fmt.Println("\nSoal 1")
	var luasLigkaran float64
	var kelilingLingkaran float64

	var phi float64 = 22 / 7
	var jariJari float64 = 35

	luasLigkaran = luas(phi, jariJari)
	kelilingLingkaran = keliling(phi, jariJari)
	fmt.Println("Luas lingkaran sebelum perubahan adalah", luasLigkaran)
	fmt.Println("Keliling lingkaran sebelum perubahan adalah", kelilingLingkaran)

	updateValue(&jariJari, 28)
	luasLigkaran = luas(phi, jariJari)
	kelilingLingkaran = keliling(phi, jariJari)
	fmt.Println("Luas lingkaran setelah perubahan", luasLigkaran)
	fmt.Println("Keliling lingkaran setelah perubahan", kelilingLingkaran)
}

func luas(phi float64, jariJari float64) float64 {
	return phi * jariJari * jariJari
}

func keliling(phi float64, jariJari float64) float64 {
	return phi * 2 * jariJari
}
func updateValue(original *float64, value float64) {
	*original = value
}

func soal2() {
	fmt.Println("\nSoal 2")
	var sentence string

	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)
}

func introduce(sentence *string, nama string, jenisKelamin string, pekerjaan string, umur string) {
	var newSentence string

	if jenisKelamin == "laki-laki" {
		newSentence = "Pak " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun "
	} else {
		newSentence = "Bu " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}
	*sentence = newSentence
}

func soal3() {
	fmt.Println("\nSoal 3")
	var buah = []string{}
	var listBuah = []string{"test", "Jeruk", "Semangka", "Mangga",
		"Strawberry", "Durian", "Manggis", "Alpukat"}
	updateToBuah(&buah, listBuah...)

	for i := 1; i < len(buah); i++ {
		fmt.Print(i, ". ", buah[i], "\n")
	}

}

func updateToBuah(slice *[]string, slice_list ...string) {
	for _, add := range slice_list {
		*slice = append(*slice, add)
	}

}

func soal4() {
	fmt.Println("\nSoal 4")
	var dataFilm = []map[string]string{}

	tambahDataFilm(&dataFilm, "LOTR", "2 jam", "action", "1999")
	tambahDataFilm(&dataFilm, "avenger", "2 jam", "action", "2019")
	tambahDataFilm(&dataFilm, "spiderman", "2 jam", "action", "2004")
	tambahDataFilm(&dataFilm, "juon", "2 jam", "horror", "2004")

	i := 1
	for _, item := range dataFilm {
		fmt.Print(i, ". ")
		fmt.Println("title :", item["title"])
		fmt.Println("  ", "duration :", item["duration"])
		fmt.Println("  ", "genre :", item["genre"])
		fmt.Println("  ", "year :", item["year"], "\n")
		i++
	}
}

func tambahDataFilm(dataFilm *[]map[string]string, nama string, durasi string, genre string, year string) {
	var data = map[string]string{}
	data["genre"] = genre
	data["duration"] = durasi
	data["year"] = year
	data["title"] = nama
	*dataFilm = append(*dataFilm, data)
}
