package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	soal1()
	fmt.Println("")
	soal2()
	fmt.Println("")
	soal3()
	soal4()
	fmt.Println("")
	soal5()
}

// Soal 1
func introduce(sentence *string, name string, gender string, occupation string, age string) {
	var referToAs string
	if gender == "laki-laki" {
		referToAs = "Pak"
	} else {
		referToAs = "Bu"
	}

	*sentence = referToAs + " " + name + " adalah seorang " + occupation + " yang berusia " + age + " tahun"
}

func soal1() {
	var sentence string

	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)
}

// Soal 2
func fruitsList(fruits *[]string) {
	buahBuahan := []string{"Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat"}
	*fruits = append(*fruits, buahBuahan...)

	for i, fruit := range *fruits {
		fmt.Printf("%d. %s\n", i+1, fruit)
	}
}

func soal2() {
	var buah = []string{}

	fruitsList(&buah)
}

// Soal 3
func tambahDataFilm(title string, runtime string, genre string, releaseYear string, data *[]map[string]string) {
	filmToAdd := map[string]string{
		"title":    title,
		"duration": runtime,
		"genre":    genre,
		"year":     releaseYear,
	}
	*data = append(*data, filmToAdd)
}

func showFilms(films *[]map[string]string) {
	for i, film := range *films {
		fmt.Printf("%d. Judul: %s\n", i+1, film["title"])
		fmt.Printf("   Durasi: %s\n", film["duration"])
		fmt.Printf("   Genre: %s\n", film["genre"])
		fmt.Printf("   Tahun: %s\n", film["year"])
		fmt.Println("")
	}
}

func soal3() {
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	showFilms(&dataFilm)
}

// Soal 4
func arrayMaker(array *[5]int, numbers ...int) {
	fmt.Println("Array Before:", *array)
	// for i := 0; i < len(numbers); i++ {
	// 	(*array)[i] = numbers[i]
	// }
	copy((*array)[:], numbers)
	fmt.Println("Array Now (Before Being Evened Out):", *array)
}

func evenizer(array *[5]int) {
	for i := 0; i < len(array); i++ {
		if (*array)[i]%2 != 0 {
			(*array)[i] *= 2
		}
	}
	fmt.Println("Array Now (After Being Evened Out):", *array)
}

func soal4() {
	var arr [5]int
	arrayMaker(&arr, 1, 2, 3, 4, 5)
	evenizer(&arr)
}

// Soal 5
type Buah struct {
	Nama       string `json:"nama"`
	Warna      string `json:"warna"`
	AdaBijinya bool   `json:"ada_bijinya"`
	Harga      int    `json:"harga"`
}

func printStruct(fruit Buah) {
	jsonData, err := json.MarshalIndent(fruit, "", " ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func soal5() {
	pineapple := Buah{
		Nama:       "Nanas",
		Warna:      "Kuning",
		AdaBijinya: false,
		Harga:      9000,
	}
	orange := Buah{
		Nama:       "Jeruk",
		Warna:      "Oranye",
		AdaBijinya: true,
		Harga:      8000,
	}
	watermelon := Buah{
		Nama:       "Semangka",
		Warna:      "Hijau & Merah",
		AdaBijinya: true,
		Harga:      10000,
	}
	banana := Buah{
		Nama:       "Pisang",
		Warna:      "Kuning",
		AdaBijinya: false,
		Harga:      5000,
	}

	printStruct(pineapple)
	printStruct(orange)
	printStruct(watermelon)
	printStruct(banana)
}
