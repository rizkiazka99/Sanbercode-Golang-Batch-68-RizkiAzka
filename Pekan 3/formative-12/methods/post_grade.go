package methods

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type NilaiMahasiswa struct {
	ID          uint   `json:"id"`
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string `json:"indeks_nilai"`
	Nilai       uint   `json:"nilai"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func idGenerator() string {
	currentTime := time.Now()
	timestamp := currentTime.Format("2006-01-02 15:04:05.000")
	charsToRemove := [...]string{" ", "-", ":", "."}
	for _, char := range charsToRemove {
		timestamp = strings.ReplaceAll(timestamp, string(char), "")
	}

	return timestamp
}

func gradeIndex(grade uint) string {
	if grade >= 80 {
		return "A"
	} else if grade >= 70 && grade < 80 {
		return "B"
	} else if grade >= 60 && grade < 70 {
		return "C"
	} else if grade >= 50 && grade < 60 {
		return "D"
	} else if grade < 50 {
		return "E"
	} else {
		return "F"
	}
}

func PostGrade(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var grade NilaiMahasiswa

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&grade); err != nil {
				log.Fatal(err)
			}

			getID := idGenerator()
			intId, _ := strconv.Atoi(getID)
			grade.ID = uint(intId)
			grade.IndeksNilai = gradeIndex(grade.Nilai)
		} else {
			getID := idGenerator()
			intId, _ := strconv.Atoi(getID)
			id := uint(intId)

			getNilai := r.PostFormValue("nilai")
			intNilai, _ := strconv.Atoi(getNilai)
			nilai := uint(intNilai)

			nama := r.PostFormValue("nama")
			mataKuliah := r.PostFormValue("mata_kuliah")
			indeksNilai := gradeIndex(nilai)

			grade = NilaiMahasiswa{
				ID:          id,
				Nilai:       nilai,
				Nama:        nama,
				MataKuliah:  mataKuliah,
				IndeksNilai: indeksNilai,
			}
		}

		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, grade)
		dataGrade, _ := json.Marshal(grade)
		w.Write(dataGrade)
		return
	} else {
		http.Error(w, "NOT FOUND", http.StatusNotFound)
		return
	}
}
