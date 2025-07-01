package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	soal1()
	soal2()
	soal3()
	soal4()
	soal5()
}

func soal1() {
	firstWord := "Bootcamp"
	secondWord := "Digital"
	thirdWord := "Skill"
	fourthWord := "Sanbercode"
	fifthWord := "Golang"

	fmt.Println("" + firstWord + " " + secondWord + " " + thirdWord + " " + fourthWord + " " + fifthWord + "")
}

func soal2() {
	halo := "Halo Dunia"

	var newText = strings.Replace(halo, "Dunia", "Golang", -1)
	fmt.Println(newText)
}

func soal3() {
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataKeduaModified := strings.Replace(kataKedua, "s", "S", 1)
	kataKetigaModified := strings.Replace(kataKetiga, "r", "R", 1)
	kataKeempatModified := strings.ToUpper(kataKeempat)

	fmt.Println("" + kataPertama + " " + kataKeduaModified + " " + kataKetigaModified + " " + kataKeempatModified + "")
}

func soal4() {
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	firstNum, _ := strconv.Atoi(angkaPertama)
	secondNum, _ := strconv.Atoi(angkaKedua)
	thirdNum, _ := strconv.Atoi(angkaKetiga)
	fourthNum, _ := strconv.Atoi(angkaKeempat)

	result := firstNum + secondNum + thirdNum + fourthNum
	fmt.Println(result)
}

func soal5() {
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjangPersegiPanjangNum, _ := strconv.Atoi(panjangPersegiPanjang)
	lebarPersegiPanjangNum, _ := strconv.Atoi(lebarPersegiPanjang)
	alasSegitigaNum, _ := strconv.Atoi(alasSegitiga)
	tinggiSegitigaNum, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	luasPersegiPanjang = panjangPersegiPanjangNum * lebarPersegiPanjangNum
	kelilingPersegiPanjang = 2 * (panjangPersegiPanjangNum + lebarPersegiPanjangNum)
	luasSegitiga = int(0.5 * float64(alasSegitigaNum) * float64(tinggiSegitigaNum))

	luasPersegiPanjangStr := strconv.Itoa(luasPersegiPanjang)
	kelilingPersegiPanjangStr := strconv.Itoa(kelilingPersegiPanjang)
	luasSegitigaStr := strconv.Itoa(luasSegitiga)

	println("Luas Persegi Panjang: " + luasPersegiPanjangStr + ", Keliling Persegi Panjang: " + kelilingPersegiPanjangStr + ", Luas Segitiga: " + luasSegitigaStr + "")

}
