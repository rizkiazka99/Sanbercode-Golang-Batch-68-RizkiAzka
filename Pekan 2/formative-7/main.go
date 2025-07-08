package main

import (
	"fmt"
)

func main() {
	soal1()
	fmt.Println("")
	soal2()
	fmt.Println("")
	soal3()
	fmt.Println("")
	soal4()
	soal5()
}

func soal1() {
	type Mahasiswa struct {
		Nama string
		NIM  string
		Usia int
	}

	student := Mahasiswa{
		Nama: "Rizki Azka Fihi Aghnia",
		NIM:  "210340003",
		Usia: 22,
	}

	fmt.Println("Nama:", student.Nama)
	fmt.Println("NIM:", student.NIM)
	fmt.Printf("Usia: %d\n", student.Usia)
}

// Soal 2
type segitiga struct {
	alas, tinggi int
}

func (segitiga segitiga) calcTriangleArea() float32 {
	area := 0.5 * float32(segitiga.alas) * float32(segitiga.tinggi)
	return area
}

type persegi struct {
	sisi int
}

func (persegi persegi) calcSquareArea() int {
	area := persegi.sisi * persegi.sisi
	return area
}

type persegiPanjang struct {
	panjang, lebar int
}

func (persegiPanjang persegiPanjang) calcRectArea() int {
	area := persegiPanjang.panjang * persegiPanjang.lebar
	return area
}

func soal2() {
	var triangle = segitiga{
		alas:   6,
		tinggi: 10,
	}
	var square = persegi{
		sisi: 15,
	}
	var rectangle = persegiPanjang{
		panjang: 20,
		lebar:   8,
	}

	fmt.Printf("Luas Segitiga: %f\n", triangle.calcTriangleArea())
	fmt.Printf("Luas Persegi: %d\n", square.calcSquareArea())
	fmt.Printf("Luas Persegi Panjang: %d\n", rectangle.calcRectArea())
}

// Soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (phone phone) addColor(smartphone *phone, colors ...string) {
	(*smartphone).colors = append((*smartphone).colors, colors...)
}

func soal3() {
	phone := phone{
		name:   "X7 Pro",
		brand:  "POCO",
		year:   2025,
		colors: []string{},
	}
	colors := []string{"Black", "Blue", "Green", "Dark Grey"}
	var colorsStr string

	phone.addColor(&phone, colors...)

	fmt.Printf("Brand: %s\n", phone.brand)
	fmt.Printf("Name: %s\n", phone.name)
	fmt.Printf("Year: %d\n", phone.year)

	for i, color := range phone.colors {
		if i != len(phone.colors)-1 {
			colorsStr += color + ", "
		} else {
			colorsStr += color
		}
	}
	fmt.Printf("Colors: %s\n", colorsStr)
}

// Soal 4
type movie struct {
	title    string
	genre    string
	duration int
	year     int
}

func tambahDataFilm(title string, runtime int, genre string, releaseYear int, dataFilm *[]movie) {
	film := movie{
		title:    title,
		genre:    genre,
		duration: runtime / 60,
		year:     releaseYear,
	}
	*dataFilm = append(*dataFilm, film)
}

func showFilms(dataFilm *[]movie) {
	for i, film := range *dataFilm {
		fmt.Printf("%d. Judul: %s\n", i+1, film.title)
		fmt.Printf("   Durasi: %d jam\n", film.duration)
		fmt.Printf("   Genre: %s\n", film.genre)
		fmt.Printf("   Tahun: %d\n", film.year)
		fmt.Println("")
	}
}

func soal4() {
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	showFilms(&dataFilm)
}

// Soal 5
type Hewan interface {
	Suara() string
}

type Kucing struct{}

func (cat Kucing) Suara() string {
	return "Meow"
}

type Anjing struct{}

func (dog Anjing) Suara() string {
	return "Woof"
}

func soal5() {
	var animal Hewan

	animal = Kucing{}
	fmt.Println("Cat goes", animal.Suara())

	animal = Anjing{}
	fmt.Println("Dog goes", animal.Suara())
}
