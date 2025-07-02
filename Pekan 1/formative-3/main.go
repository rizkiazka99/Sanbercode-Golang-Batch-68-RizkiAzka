package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	soal1()

	var nilaiJohn = 80
	var nilaiDoe = 50

	soal2(nilaiJohn)
	soal2(nilaiDoe)

	var tanggal = 15
	var bulan = 6
	var tahun = 1999

	soal3(tanggal, bulan, tahun)

	soal4(1999)

	soal5()
}

func soal1() {
	kalimat := "halo halo bandung"
	angka := 2021

	newKalimat := strings.Replace(kalimat, "halo", "Hi", 2)
	angkaStr := strconv.Itoa(angka)

	fmt.Println(`"` + newKalimat + `" - ` + angkaStr + ``)
}

func soal2(grade int) {
	if grade >= 80 {
		fmt.Println("A")
	} else if grade >= 70 && grade < 80 {
		fmt.Println("B")
	} else if grade >= 60 && grade < 70 {
		fmt.Println("C")
	} else if grade >= 50 && grade < 60 {
		fmt.Println("D")
	} else if grade < 50 {
		fmt.Println("E")
	}
}

func soal3(day int, month int, year int) {
	strDay := strconv.Itoa(day)
	var strMonth string
	strYear := strconv.Itoa(year)

	switch month {
	case 1:
		strMonth = "Januari"
	case 2:
		strMonth = "Februari"
	case 3:
		strMonth = "Maret"
	case 4:
		strMonth = "April"
	case 5:
		strMonth = "Mei"
	case 6:
		strMonth = "Juni"
	case 7:
		strMonth = "Juli"
	case 8:
		strMonth = "Agustus"
	case 9:
		strMonth = "September"
	case 10:
		strMonth = "Oktober"
	case 11:
		strMonth = "November"
	case 12:
		strMonth = "Desember"
	default:
		strMonth = "UNDEFINED"
	}

	fmt.Println("" + strDay + " " + strMonth + " " + strYear + "")
}

func soal4(birthYear int) {
	if birthYear >= 1944 && birthYear <= 1964 {
		fmt.Println("Baby Boomer")
	} else if birthYear >= 1965 && birthYear <= 1979 {
		fmt.Println("Gen X")
	} else if birthYear >= 1980 && birthYear <= 1994 {
		fmt.Println("Gen Y (Millenials)")
	} else if birthYear >= 1995 && birthYear <= 2015 {
		fmt.Println("Gen Z")
	} else if birthYear >= 2016 {
		fmt.Println("Gen Alpha")
	}
}

func soal5() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%2 != 0 {
			fmt.Println("" + strconv.Itoa(i) + " - I Love Coding")
		} else {
			if i%2 == 0 {
				fmt.Println("" + strconv.Itoa(i) + " - Berkualitas")
			} else if i%2 != 0 {
				fmt.Println("" + strconv.Itoa(i) + " - Santai")
			}
		}
	}
}
