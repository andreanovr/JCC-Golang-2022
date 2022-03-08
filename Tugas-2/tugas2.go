package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Soal 1
	var a = "Jabar "
	var b = "Coding "
	var c = "Camp "
	var d = "Golang "
	var e = "2022"
	fmt.Println("\nSoal 1")
	fmt.Println(a + b + c + d + e + "\n")

	// Soal 2
	var text = "Halo Dunia"
	var find = "Dunia"
	var replaceWith = "Golang"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println("Soal 2")
	fmt.Println(newText1 + "\n")

	// Soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	//Converting
	var kataSatu = kataPertama + " "
	var kataDua = strings.Title(kataKedua) + " "
	var fin = "r"
	var rep = "R"
	var kataTiga = strings.Replace(kataKetiga, fin, rep, 1) + " "
	var kataEmpat = strings.ToUpper(kataKeempat + " ")

	//Print
	fmt.Println("Soal 3")
	fmt.Println(kataSatu + kataDua + kataTiga + kataEmpat + "\n")

	// Soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	//Converting
	num, _ := strconv.ParseInt(angkaPertama, 10, 64)
	num1, _ := strconv.ParseInt(angkaKedua, 10, 64)
	num2, _ := strconv.ParseInt(angkaKetiga, 10, 64)
	num3, _ := strconv.ParseInt(angkaKeempat, 10, 64)

	//print
	fmt.Println("Soal 4")
	fmt.Println(num + num1 + num2 + num3)

	// Soal 5
	kalimat := "halo halo bandung"
	angka := 2022

	//Converting
	finding := "halo"
	replacing := "Hi"

	kalimat = strings.Replace(kalimat, finding, replacing, 2)
	str_angka := strconv.Itoa(angka)

	//Print
	fmt.Println("\nSoal 5")
	fmt.Println(`"` + kalimat + `"` + " - " + str_angka + "\n")

}
