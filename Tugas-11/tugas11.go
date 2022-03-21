package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
)

// Soal 1
func tambahDataHP(s *[]string) {
	for i, hape := range *s {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(strconv.Itoa(i+1)+".", hape)
	}

}

// Soal 2
func getMovies(z chan string, u ...string) {
	fmt.Println("List Movies")
	for ab, xa := range u {

		z <- (strconv.Itoa(ab+1) + ". " + xa)
	}
	close(z)
}

// soal 3
func luasLingkaran(jari2 chan float64, jari3 float64) {

	hasil := math.Pi * (math.Pow(jari3, 2))
	jari2 <- math.Ceil(hasil)

}
func kelilingLingkaran(jari2 chan float64, jari3 float64) {

	hasil := 2 * math.Pi * jari3
	jari2 <- math.Ceil(hasil)

}

func volumeTabung(jari2 chan float64, jari3 float64) {

	tinggi := 10
	hasil := math.Pi * (math.Pow(jari3, 2)) * float64(tinggi)
	jari2 <- math.Ceil(hasil)

}

func kelilingPersegiPanjang(p, l int, cha chan int) {
	hasil := 2 * (p + l)
	cha <- hasil
}

func luasPersegiPanjang(p, l int, cha chan int) {
	hasil := p * l
	cha <- hasil
}

func volumeBalok(p, l, t int, cha chan int) {
	hasil := p * l * t
	cha <- hasil
}

func main() {

	// Soal 1
	fmt.Println("\nSoal 1")
	var wg sync.WaitGroup
	wg.Add(1)
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	go func() {
		tambahDataHP(&phones)
		wg.Done()
	}()

	wg.Wait()

	// Soal 2
	fmt.Println("\nSoal 2")
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	// Soal 3
	fmt.Println("\nSoal 3")
	jari := make(chan float64)
	jari2 := make(chan float64)
	jari3 := make(chan float64)
	kll := make(chan float64)
	kll2 := make(chan float64)
	kll3 := make(chan float64)
	vol := make(chan float64)
	vol2 := make(chan float64)
	vol3 := make(chan float64)

	go func() {
		luasLingkaran(jari, 8)
		luasLingkaran(jari2, 14)
		luasLingkaran(jari3, 20)
		kelilingLingkaran(kll, 8)
		kelilingLingkaran(kll2, 14)
		kelilingLingkaran(kll3, 20)
		kelilingLingkaran(vol, 8)
		kelilingLingkaran(vol2, 14)
		kelilingLingkaran(vol3, 20)

	}()

	fmt.Println("Hasil luas lingkaran 1 :", <-jari)
	fmt.Println("Hasil luas lingkaran 2 :", <-jari2)
	fmt.Println("Hasil luas lingkaran 3 :", <-jari3)
	fmt.Println("Hasil keliling lingkaran 1 :", <-kll)
	fmt.Println("Hasil keliling lingkaran 2 :", <-kll2)
	fmt.Println("Hasil keliling lingkaran 3 :", <-kll3)
	fmt.Println("Hasil volume tabung 1 :", <-vol)
	fmt.Println("Hasil volume tabung 2 :", <-vol2)
	fmt.Println("Hasil volume tabung 3 :", <-vol3)

	// Soal 4
	fmt.Println("\nSoal 4")
	var panjang = 21
	var lebar = 7
	var tinggi = 10
	var cha1 = make(chan int)
	go luasPersegiPanjang(panjang, lebar, cha1)

	var cha2 = make(chan int)
	go kelilingPersegiPanjang(panjang, lebar, cha2)

	var cha3 = make(chan int)
	go volumeBalok(panjang, lebar, tinggi, cha3)

	for i := 0; i < 3; i++ {
		select {
		case luas := <-cha1:
			fmt.Printf("Luas Persegi : %d\n", luas)
		case keliling := <-cha2:
			fmt.Printf("keliling Persegi : %d\n", keliling)
		case volumeB := <-cha3:
			fmt.Printf("Volume Balok : %d\n", volumeB)
		}
	}
}
