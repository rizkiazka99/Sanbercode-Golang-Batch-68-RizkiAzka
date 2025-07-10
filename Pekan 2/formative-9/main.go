package main

import (
	"fmt"
	"formative-9/formative8"
)

func main() {
	soal1()
	fmt.Println("")
	soal2()
	fmt.Println("")
	soal3()
	fmt.Println("")
	formative8.Soal4()
	fmt.Println("")
	formative8.Soal5()

}

func soal1() {
	var datar formative8.HitungBangunDatar
	var ruang formative8.HitungBangunRuang

	datar = formative8.SegitigaSamaSisi{
		Alas:   12,
		Tinggi: 15,
	}
	fmt.Println("Luas Segitiga:", datar.Luas())
	fmt.Println("Keliling Segitiga:", datar.Keliling())
	fmt.Println("")

	datar = formative8.PersegiPanjang{
		Panjang: 16,
		Lebar:   8,
	}
	fmt.Println("Luas Persegi Panjang:", datar.Luas())
	fmt.Println("Keliling Persegi Panjang:", datar.Keliling())
	fmt.Println("")

	ruang = formative8.Tabung{
		JariJari: 9,
		Tinggi:   12,
	}
	fmt.Println("Volume Tabung:", ruang.Volume())
	fmt.Println("Luas Permukaan Tabung:", ruang.LuasPermukaan())
	fmt.Println("")

	ruang = formative8.Balok{
		Panjang: 10,
		Lebar:   8,
		Tinggi:  12,
	}
	fmt.Println("Volume Balok:", ruang.Volume())
	fmt.Println("Luas Permukaan Balok:", ruang.LuasPermukaan())
}

func soal2() {
	var smartphone formative8.PropertyChanger

	smartphone = formative8.Phone{
		Name:   "Galaxy Note 20",
		Brand:  "Samsung",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}
	smartphone.ShowPhoneDetails()
}

func soal3() {
	fmt.Println(formative8.LuasPersegi(4, true))
	fmt.Println(formative8.LuasPersegi(8, false))
	fmt.Println(formative8.LuasPersegi(0, true))
	fmt.Println(formative8.LuasPersegi(0, false))
}
