package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"strconv"
	"time"
)

type lingkaran struct {
	jariJari float64
}

type hitungBangunDatar interface {
	luas() float64
	keliling() float64
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
	// Soal 5
	soal5()
	// Soal 6
	soal6()
}

func printJCC(kalimat string, tahun int) {
	fmt.Println(kalimat, tahun)
}

func soal1() {
	fmt.Println("\nSoal 1")
	golang := "Golang Backend Development"
	year := 2021
	defer printJCC(golang, year)
}

func soal2() {
	fmt.Println("\nSoal 2")
	defer kelilingSegitigaSamaSisi(4, true)
	defer kelilingSegitigaSamaSisi(8, false)
	defer kelilingSegitigaSamaSisi(0, true)
	defer kelilingSegitigaSamaSisi(0, false)
}

func kelilingSegitigaSamaSisi(sisi int, stats bool) error {
	var infoKeliling string
	s := strconv.Itoa(sisi)
	kelilingSegitga := 3 * sisi
	l := strconv.Itoa(kelilingSegitga)
	if sisi > 0 && stats == true {
		infoKeliling = "keliling segitiga sama sisinya dengan sisi " + s + "cm adalah " + l + " cm"
		fmt.Println(infoKeliling)
	} else if sisi > 0 && stats == false {
		infoKeliling = l
		fmt.Println(infoKeliling)
	} else if sisi == 0 && stats == true {
		return errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	} else if sisi == 0 && stats == false {
		return errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}
	return nil
}

func recoveryFunction() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
func tambahAngka(angka1 int, angka2 *int) {
	*angka2 = angka1 + *angka2
	fmt.Println(*angka2)
}

func cetakAngka(angka *int) {
	fmt.Println(*angka)
}

func soal3() {
	fmt.Println("\nSoal 3")
	angka := 1
	defer cetakAngka(&angka)
	defer tambahAngka(7, &angka)
	defer tambahAngka(6, &angka)
	defer tambahAngka(-1, &angka)
	defer tambahAngka(9, &angka)
	fmt.Println("Berhasil menjumlahkan angka:")
}

func soal4() {
	fmt.Println("\nSoal 4")
	var phones = []string{}
	var phoneName = []string{"Xiaomi", "Asus", "Iphone", "Samsung",
		"Oppo", "Realme", "Vivo"}
	addPhones(&phones, phoneName...)

	for i := 1; i < len(phones); i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(i, ".", phones[i])
	}
}

func addPhones(phones *[]string, name ...string) {
	for _, phone := range name {
		*phones = append(*phones, phone)
	}
}

func soal5() {
	fmt.Println("\nSoal 5")
	var lingkaran1, lingkaran2, lingkaran3 hitungBangunDatar
	lingkaran1 = lingkaran{7}
	lingkaran2 = lingkaran{10}
	lingkaran3 = lingkaran{15}

	fmt.Println("Luas lingkaran 1 adalah: ", lingkaran1.luas())
	fmt.Println("Keliling lingkaran 1 adalah: ", lingkaran1.keliling())
	fmt.Println("Luas lingkaran 2 adalah: ", lingkaran2.luas())
	fmt.Println("Keliling lingkaran 2 adalah: ", lingkaran2.keliling())
	fmt.Println("Luas lingkaran 3 adalah: ", lingkaran3.luas())
	fmt.Println("Keliling lingkaran 3 adalah: ", lingkaran3.keliling())
}

func (l lingkaran) luas() float64 {
	return math.Round(math.Pi * math.Pow(l.jariJari, 2))
}
func (l lingkaran) keliling() float64 {
	return math.Round(math.Pi * (l.jariJari + l.jariJari))
}
func soal6() {
	fmt.Println("\nSoal 6")
	var panjang = flag.Int64("panjang", 20, " persegiPanjang")
	var lebar = flag.Int64("lebar", 8, " persegiPanjang")
	flag.Parse()
	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang(*panjang, *lebar))
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang(*panjang, *lebar))
}

func luasPersegiPanjang(panjang int64, lebar int64) int64 {
	return panjang * lebar
}
func kelilingPersegiPanjang(panjang int64, lebar int64) int64 {
	return (2 * panjang) + (2 * lebar)
}
