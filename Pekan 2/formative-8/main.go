package main

import (
	"fmt"
	"formative-8/matematika"
	"math"
	"strconv"
	"strings"
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
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return int(0.5 * float64(s.alas) * float64(s.tinggi))
}

func (s segitigaSamaSisi) keliling() int {
	var sisi float64 = (2 * float64(s.tinggi)) / math.Sqrt(3)
	return int(math.Round(sisi) * 3)
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return (2 * p.panjang) + (2 * p.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang) * float64(b.lebar) * float64(b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return float64(2 * ((b.panjang * b.lebar) + (b.panjang*b.tinggi)*(b.lebar*b.tinggi)))
}

func soal1() {
	var datar hitungBangunDatar
	var ruang hitungBangunRuang

	datar = segitigaSamaSisi{
		alas:   12,
		tinggi: 15,
	}
	fmt.Println("Luas Segitiga:", datar.luas())
	fmt.Println("Keliling Segitiga:", datar.keliling())
	fmt.Println("")

	datar = persegiPanjang{
		panjang: 16,
		lebar:   8,
	}
	fmt.Println("Luas Persegi Panjang:", datar.luas())
	fmt.Println("Keliling Persegi Panjang:", datar.keliling())
	fmt.Println("")

	ruang = tabung{
		jariJari: 9,
		tinggi:   12,
	}
	fmt.Println("Volume Tabung:", ruang.volume())
	fmt.Println("Luas Permukaan Tabung:", ruang.luasPermukaan())
	fmt.Println("")

	ruang = balok{
		panjang: 10,
		lebar:   8,
		tinggi:  12,
	}
	fmt.Println("Volume Balok:", ruang.volume())
	fmt.Println("Luas Permukaan Balok:", ruang.luasPermukaan())
}

// Soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type propertyChanger interface {
	showPhoneDetails()
}

func (phone phone) showPhoneDetails() {
	var colorsStr string

	fmt.Printf("Brand : %s\n", phone.brand)
	fmt.Printf("Name : %s\n", phone.name)
	fmt.Printf("Year : %d\n", phone.year)

	for i, color := range phone.colors {
		if i != len(phone.colors)-1 {
			colorsStr += color + ", "
		} else {
			colorsStr += color
		}
	}
	fmt.Printf("Colors : %s\n", colorsStr)
}

func soal2() {
	var smartphone propertyChanger

	smartphone = phone{
		name:   "Galaxy Note 20",
		brand:  "Samsung",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}
	smartphone.showPhoneDetails()
}

// Soal 3
func luasPersegi(sisi int, returnMessage bool) (result interface{}) {
	luas := sisi * sisi
	luasStr := strconv.Itoa(luas)
	sisiStr := strconv.Itoa(sisi)

	if returnMessage {
		if sisi > 0 {
			result = "Luas persegi dengan sisi " + sisiStr + " cm adalah " + luasStr + " cm"
		} else {
			result = "Maaf anda belum menginput sisi dari persegi"
		}
	} else {
		if sisi > 0 {
			result = luasStr
		} else {
			result = nil
		}
	}

	return
}

func soal3() {
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))
}

// Soal 4
func intsToStrings(ints []int) []string {
	result := make([]string, len(ints))

	for i, v := range ints {
		result[i] = strconv.Itoa(v)
	}
	return result
}

func soal4() {
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	firstInts := kumpulanAngkaPertama.([]int)
	secondInts := kumpulanAngkaKedua.([]int)

	firstNumbers := strings.Join(intsToStrings(firstInts), " + ")
	secondNumbers := strings.Join(intsToStrings(secondInts), " + ")
	numbers := firstNumbers + " + " + secondNumbers

	var result int = 0
	for i := 0; i < len(kumpulanAngkaPertama.([]int)); i++ {
		result += kumpulanAngkaPertama.([]int)[i]
		result += kumpulanAngkaKedua.([]int)[i]
	}
	resultStr := strconv.Itoa(result)

	fmt.Println(prefix.(string) + numbers + " = " + resultStr)
}

func soal5() {
	a := 5
	b := 3

	hasilTambah := matematika.Tambah(a, b)
	hasilKali := matematika.Kali(a, b)

	fmt.Printf("Hasil Tambah %d + %d = %d\n", a, b, hasilTambah)
	fmt.Printf("Hasil Kali %d x %d = %d\n", a, b, hasilKali)
}
