package main

import (
	"fmt"
)

type buah struct {
	nama, warna string
	adaBijinya  bool
	harga       int
}

type segitiga struct {
	alas, tinggi int
}
type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type movie struct {
	title, genre   string
	year, duration int
}

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
	var nanas = buah{"Nanas", "Kuning", false, 9000}
	var jeruk = buah{"Jeruk", "Oranye", true, 8000}
	var semangka = buah{"Semangka", "Hijau & Merah", true, 10000}
	var pisang = buah{"Pisang", "Kuning", false, 5000}

	fmt.Println("Nama |", "Warna |", "Ada Bijinya |", "Harga")

	if nanas.adaBijinya == true {
		adaBiji := "Ada"
		fmt.Println(nanas.nama, "|", nanas.warna, "|", adaBiji, "|", nanas.harga)
	} else {
		adaBiji := "Tidak"
		fmt.Println(nanas.nama, "|", nanas.warna, "|", adaBiji, "|", nanas.harga)
	}
	if jeruk.adaBijinya == true {
		adaBiji := "Ada"
		fmt.Println(jeruk.nama, "|", jeruk.warna, "|", adaBiji, "|", jeruk.harga)
	} else {
		adaBiji := "Tidak"
		fmt.Println(jeruk.nama, "|", jeruk.warna, "|", adaBiji, "|", jeruk.harga)
	}
	if pisang.adaBijinya == true {
		adaBiji := "Ada"
		fmt.Println(pisang.nama, "|", pisang.warna, "|", adaBiji, "|", pisang.harga)
	} else {
		adaBiji := "Tidak"
		fmt.Println(pisang.nama, "|", pisang.warna, "|", adaBiji, "|", pisang.harga)
	}
	if semangka.adaBijinya == true {
		adaBiji := "Ada"
		fmt.Println(semangka.nama, "|", semangka.warna, "|", adaBiji, "|", semangka.harga)
	} else {
		adaBiji := "Tidak"
		fmt.Println(semangka.nama, "|", semangka.warna, "|", adaBiji, "|", semangka.harga)
	}
}

func soal2() {
	fmt.Println("\nSoal 2")
	var persegi = persegi{6}
	var persegiPanjang = persegiPanjang{6, 10}
	var segitiga = segitiga{10, 4}

	persegi.luasPersegi()
	persegiPanjang.luasPersegiPanjang()
	segitiga.luasSegitiga()
}

func (p persegi) luasPersegi() {
	luas := p.sisi * p.sisi
	fmt.Println("Luas persegi adalah :", luas)
}
func (s segitiga) luasSegitiga() {
	luas := s.alas * s.tinggi / 2
	fmt.Println("Luas segitiga adalah :", luas)
}
func (pp persegiPanjang) luasPersegiPanjang() {
	luas := pp.panjang * pp.lebar
	fmt.Println("Luas persegi panjang adalah :", luas)
}
func soal3() {
	fmt.Println("\nSoal 3")
	var color []string
	var colorList = []string{"Graphite,", "Gold,", "Silver,", "Black,", "Sierra Blue"}
	Iphone := phone{"Iphone 13 Pro Max,", "Iphone,", 2022, color}
	color = phone.addColors(Iphone, colorList...)
	Iphone.colors = color
	fmt.Println(Iphone)
}

func (p phone) addColors(warna ...string) []string {
	var colorPalette []string
	for _, item := range warna {
		colorPalette = append(colorPalette, item)
	}
	return colorPalette
}

func soal4() {
	fmt.Println("\nSoal 4")
	var dataFilm = []map[string]string{}
	var lotr, avenger, spiderman, juon movie

	movie.tambahDataFilm(lotr, &dataFilm, "LOTR", "2 jam", "action", "1999")
	movie.tambahDataFilm(avenger, &dataFilm, "avenger", "2 jam", "action", "2019")
	movie.tambahDataFilm(spiderman, &dataFilm, "spiderman", "2 jam", "action", "2004")
	movie.tambahDataFilm(juon, &dataFilm, "juon", "2 jam", "horror", "2004")

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

func (m movie) tambahDataFilm(dataFilm *[]map[string]string, nama string, durasi string, genre string, year string) {
	var data = map[string]string{}
	data["genre"] = genre
	data["duration"] = durasi
	data["year"] = year
	data["title"] = nama
	*dataFilm = append(*dataFilm, data)
}
