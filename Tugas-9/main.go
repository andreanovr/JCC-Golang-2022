package main

import (
	. "Tugas-9/library"
	"fmt"
	"strconv"
)

func main() {
	// Soal 1
	fmt.Println("\nSoal 1")
	var bangunDatar HitungBangunDatar
	var bangunRuang HitungBangunRuang
	bangunDatar = PersegiPanjang{10, 5}
	fmt.Println("\nPersegi Panjang")
	fmt.Println("Luas :", bangunDatar.Luas())
	fmt.Println("keliling:", bangunDatar.Keliling())

	bangunDatar = SegitigaSamaSisi{14, 7}
	fmt.Println("\nSegitiga Sama Sisi")
	fmt.Println("luas :", bangunDatar.Luas())
	fmt.Println("keliling :", bangunDatar.Keliling())

	bangunRuang = Tabung{14.7, 7}
	fmt.Println("\nTabung")
	fmt.Println("volume :", bangunRuang.Volume())
	fmt.Println("luasPermukaan :", bangunRuang.LuasPermukaan())

	bangunRuang = Balok{9, 10, 11}
	fmt.Println("\nBalok")
	fmt.Println("volume :", bangunRuang.Volume())
	fmt.Println("luasPermukaan :", bangunRuang.LuasPermukaan())

	// Soal 2
	fmt.Println("\nSoal 2")
	var color []string
	var color_list = []string{"Mystic Bronze", "Mystic White", "Mystic Black"}
	xiaomi := Phone{"Samsung Galaxy Note 20", "Samsung Galaxy Note 20", 2020, color}
	color = Phone.AddColors(xiaomi, color_list...)
	xiaomi.Colors = color
	xiaomi.Display(xiaomi)

	// Soal 3
	fmt.Println("\nSoal 3")
	fmt.Println(LuasPersegi(4, true))
	fmt.Println(LuasPersegi(8, false))
	fmt.Println(LuasPersegi(0, true))
	fmt.Println(LuasPersegi(0, false))

	// Soal 4
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
