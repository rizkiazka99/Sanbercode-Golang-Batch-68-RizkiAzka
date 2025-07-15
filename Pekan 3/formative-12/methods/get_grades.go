package methods

import (
	"encoding/json"
	"net/http"
)

func GetGrades(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilai, err := json.Marshal(nilaiNilaiMahasiswa)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)
		return
	} else {
		http.Error(w, "ERROR...", http.StatusNotFound)
	}
}
