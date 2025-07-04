package main

import (
	"fmt"
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

func soal1() {
	var shape string
	for i := 0; i < 7; i++ {
		shape += "#"
		fmt.Println(shape)
	}
}

func soal2() {
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	sentence := kalimat[2:]

	fmt.Println(sentence)
}

func soal3() {
	var sayuran = []string{}
	var vegetables = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i, item := range vegetables {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

func soal4() {
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	satuan["Volume Balok"] = satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]

	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)
	}
}

// Soal 5
func luasPersegiPanjang(panjang int, lebar int) string {
	result := panjang * lebar
	luas := strconv.Itoa(result)

	return luas
}

func kelilingPersegiPanjang(panjang int, lebar int) string {
	result := 2 * (panjang + lebar)
	keliling := strconv.Itoa(result)

	return keliling
}

func volumeBalok(panjang int, lebar int, tinggi int) string {
	result := panjang * lebar * tinggi
	volume := strconv.Itoa(result)

	return volume
}

func soal5() {
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)
}
