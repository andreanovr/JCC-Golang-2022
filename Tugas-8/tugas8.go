package main

import (
	"fmt"
	"math"
	"strconv"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi float64
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type tampilDetailPhone interface {
	display() string
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
	var bangunDatar hitungBangunDatar
	var bangunRuang hitungBangunRuang
	bangunDatar = persegiPanjang{10, 5}
	fmt.Println("\nPersegi Panjang")
	fmt.Println("Luas :", bangunDatar.luas())
	fmt.Println("keliling:", bangunDatar.keliling())

	bangunDatar = segitigaSamaSisi{14, 7}
	fmt.Println("\nSegitiga Sama Sisi")
	fmt.Println("luas :", bangunDatar.luas())
	fmt.Println("keliling :", bangunDatar.keliling())

	bangunRuang = tabung{14.7, 7}
	fmt.Println("\nTabung")
	fmt.Println("volume :", bangunRuang.volume())
	fmt.Println("luasPermukaan :", bangunRuang.luasPermukaan())

	bangunRuang = balok{9, 10, 11}
	fmt.Println("\nBalok")
	fmt.Println("volume :", bangunRuang.volume())
	fmt.Println("luasPermukaan :", bangunRuang.luasPermukaan())
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2*p.panjang + 2*p.lebar
}

func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * 2 * t.tinggi
}
func (t tabung) luasPermukaan() float64 {
	return math.Pi * 2 * t.jariJari * t.tinggi
}

func (b balok) volume() float64 {
	return b.panjang * b.lebar * b.tinggi
}

func (b balok) luasPermukaan() float64 {
	return 2*b.panjang*b.lebar + b.panjang*b.tinggi + b.lebar*b.tinggi
}

func soal2() {
	fmt.Println("\nSoal 2")
	var color []string
	var color_list = []string{"Mystic Bronze", "Mystic White", "Mystic Black"}
	xiaomi := phone{"Samsung Galaxy Note 20", "Samsung Galaxy Note 20", 2020, color}
	color = phone.addColors(xiaomi, color_list...)
	xiaomi.colors = color
	xiaomi.display(xiaomi)

}

func (p phone) addColors(warna ...string) []string {
	var color_palette []string
	for _, item := range warna {
		color_palette = append(color_palette, item)
	}
	return color_palette
}

func (p phone) display(merk phone) {
	s := strconv.Itoa(merk.year)
	fmt.Println("name :", merk.name)
	fmt.Println("brand :", merk.brand)
	fmt.Println("year :", s)
	fmt.Println("colors :", merk.colors[0:])
}

func soal3() {
	fmt.Println("\nSoal 3")
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))
}

func luasPersegi(sisi int, stats bool) interface{} {
	var infoLuas interface{}
	s := strconv.Itoa(sisi)
	luasPersegi := sisi * sisi
	l := strconv.Itoa(luasPersegi)
	if sisi > 0 && stats == true {
		infoLuas = "luas persegi dengan sisi " + s + " adalah " + l
	} else if sisi > 0 && stats == false {
		infoLuas = l
	} else if sisi == 0 && stats == true {
		infoLuas = "Maaf anda belum menginput sisi dari persegi"
	} else if sisi == 0 && stats == false {
		infoLuas = nil
	}
	return infoLuas
}

func soal4() {
	fmt.Println("\nSoal 4")
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	kalimat := "" + prefix.(string)
	kumpulanAngka := []int{}
	kumpulanAngka = append(kumpulanAngka, kumpulanAngkaPertama.([]int)[0], kumpulanAngkaPertama.([]int)[1])
	kumpulanAngka = append(kumpulanAngka, kumpulanAngkaKedua.([]int)[0], kumpulanAngkaKedua.([]int)[1])

	jumlah := 0
	for index, item := range kumpulanAngka {
		if index == 0 {
			kalimat += strconv.Itoa(item)
		} else {
			kalimat += "+" + strconv.Itoa(item)
		}
		jumlah += item
	}

	kalimat += "=" + strconv.Itoa(jumlah)
	fmt.Println(kalimat)
}
