package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
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
func deferTest(kalimat string, tahun int) {
	fmt.Println(kalimat + " " + strconv.Itoa(tahun))
}

func soal1() {
	defer deferTest("Golang Backend Development", 2021)
	fmt.Println("Initializing...")
}

// Soal 2
func kelilingSegitigaSamaSisi(sisi int, showMessage bool) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("Maaf anda belum menginput sisi dari segitiga sama sisi. Panic: %v", r))
		}
	}()

	keliling := 3 * sisi
	sisiStr := strconv.Itoa(sisi)
	kelilingStr := strconv.Itoa(keliling)

	if sisi > 0 {
		if showMessage {
			return "Keliling segitiga sama sisinya dengan sisi " + sisiStr + " cm adalah " + kelilingStr + " cm", nil
		} else {
			return sisiStr, nil
		}
	} else {
		if showMessage {
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			panic("Input sisi bernilai 0")
		}
	}
}

func soal2() {
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
}

// Soal 3
func tambahAngka(addedBy int, baseNumber *int) {
	*baseNumber += addedBy
}

func cetakAngka(angka *int) {
	fmt.Printf("Total angka: %d\n", *angka)
}

func soal3() {
	angka := 1

	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)
}

// Soal 4
func addPhoneManufacturer(phones *[]string, manufacturers ...string) {
	*phones = append(*phones, manufacturers...)
}

func showPhoneManufacturers(phones *[]string) {
	for i, manufacturer := range *phones {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d. %s\n", i+1, manufacturer)
	}
}

func soal4() {
	var phones = []string{}
	var manufacturers = []string{"Xiaomi", "Asus", "iPhone", "Samsung", "Oppo", "Realme", "Vivo"}

	addPhoneManufacturer(&phones, manufacturers...)
	showPhoneManufacturers(&phones)
}

// Soal 5
func sortPhoneManufacturers(phones ...string) {
	sort.Strings(phones)

	var wg sync.WaitGroup
	wg.Add(len(phones))

	ch := make(chan string)

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		i := 1
		for phone := range ch {
			<-ticker.C
			fmt.Printf("%d. %s\n", i, phone)
			i++
			wg.Done()
		}
	}()

	go func() {
		for _, phone := range phones {
			ch <- phone
		}
		close(ch)
	}()

	wg.Wait()
}

func soal5() {
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	sortPhoneManufacturers(phones...)
}
