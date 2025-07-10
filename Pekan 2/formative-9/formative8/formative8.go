package formative8

import (
	"fmt"
	"formative-9/formative8/matematika"
	"math"
	"strconv"
	"strings"
)

// Soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return int(0.5 * float64(s.Alas) * float64(s.Tinggi))
}

func (s SegitigaSamaSisi) Keliling() int {
	var sisi float64 = (2 * float64(s.Tinggi)) / math.Sqrt(3)
	return int(math.Round(sisi) * 3)
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return (2 * p.Panjang) + (2 * p.Lebar)
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang) * float64(b.Lebar) * float64(b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return float64(2 * ((b.Panjang * b.Lebar) + (b.Panjang*b.Tinggi)*(b.Lebar*b.Tinggi)))
}

// Soal 2
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type PropertyChanger interface {
	ShowPhoneDetails()
}

func (phone Phone) ShowPhoneDetails() {
	var colorsStr string

	fmt.Printf("Brand : %s\n", phone.Brand)
	fmt.Printf("Name : %s\n", phone.Name)
	fmt.Printf("Year : %d\n", phone.Year)

	for i, color := range phone.Colors {
		if i != len(phone.Colors)-1 {
			colorsStr += color + ", "
		} else {
			colorsStr += color
		}
	}
	fmt.Printf("Colors : %s\n", colorsStr)
}

// Soal 3
func LuasPersegi(sisi int, returnMessage bool) (result interface{}) {
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

// Soal 4
func intsToStrings(ints []int) []string {
	result := make([]string, len(ints))

	for i, v := range ints {
		result[i] = strconv.Itoa(v)
	}
	return result
}

func Soal4() {
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

func Soal5() {
	a := 5
	b := 3

	hasilTambah := matematika.Tambah(a, b)
	hasilKali := matematika.Kali(a, b)

	fmt.Printf("Hasil Tambah %d + %d = %d\n", a, b, hasilTambah)
	fmt.Printf("Hasil Kali %d x %d = %d\n", a, b, hasilKali)
}
