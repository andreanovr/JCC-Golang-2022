package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Soal 1
	fmt.Println("\nSoal 1")
	for a := 1; a < 21; a++ {
		var num = ""
		switch {
		case a%2 == 1:
			if a%3 == 0 {
				num = "I Love Coding"
			} else {
				num = "JCC"
			}
		default:
			num = "Candradimuka"
		}
		fmt.Println(a, "-", num)
	}

	// Soal 2
	fmt.Println("\nSoal 2")
	for b := 0; b < 7; b++ {
		fmt.Println(strings.Repeat("#", b))
	}

	// Soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println("\nSoal 3")
	fmt.Println(kalimat[2:])

	// Soal 4
	fmt.Println("\nSoal 4")
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")
	for c, d := range sayuran {
		fmt.Println(strconv.Itoa(c+1) + ". " + d)
	}

	// Soal 5
	fmt.Println("\nSoal 5")
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	satuan["Volume Balok"] = satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	for f := range satuan {
		fmt.Println(f, "=", satuan[f])
	}
}
