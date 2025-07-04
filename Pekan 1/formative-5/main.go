package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	soal1()
	fmt.Println("")
	soal2()
	fmt.Println("")
	soal3()
	fmt.Println("")
	soal4()
	fmt.Println("")
	soal5()
}

// Soal 1
func introduce(name string, gender string, occupation string, age string) (sentence string) {
	var referToAs string
	if gender == "laki-laki" {
		referToAs = "Pak"
	} else {
		referToAs = "Bu"
	}

	sentence = "" + referToAs + " " + name + " adalah seorang " + occupation + " yang berusia " + age + " tahun"
	return
}

func soal1() {
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)
}

// Soal 2
func buahFavorit(name string, fruits ...string) string {
	var favorites string
	for i, fruit := range fruits {
		if i != len(fruits)-1 {
			favorites += `"` + fruit + `", `
		} else {
			favorites += `"` + fruit + `"`
		}
	}
	sentence := "Halo, nama saya " + name + " dan buah favorit saya adalah " + favorites + ""

	return sentence
}

func soal2() {
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)
}

// Soal 3
func soal3() {
	var dataFilm = []map[string]string{}

	tambahDataFilm := func(title string, runtime string, genre string, releaseYear string) {
		dataFilm = append(dataFilm, map[string]string{
			"title": title,
			"jam":   runtime,
			"tahun": releaseYear,
			"genre": genre,
		})
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

// Soal 4
func factorial(value int) {
	var result int = value
	for i := value; i > 0; i-- {
		if i == value {
			result = value
		} else {
			result *= i
		}
	}

	fmt.Println(strconv.Itoa(result))
}

func soal4() {
	factorial(5)
	factorial(7)
}

// Soal 5
func calculateCircle(luas *float64, keliling *float64, jariJari *float64) {
	*luas = math.Pi * (*jariJari) * (*jariJari)
	*keliling = 2 * math.Pi * (*jariJari)
}

func soal5() {
	var luasLingkaran float64
	var kelilingLingkaran float64
	var jariJari float64 = 7

	calculateCircle(&luasLingkaran, &kelilingLingkaran, &jariJari)

	fmt.Printf("Luas Lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran: %.2f\n", kelilingLingkaran)
}
