package Person

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Person struct {
	Nama   string `json:"Nama"`
	NIM    string `json:"NIM"`
	Alamat string `json:"Alamat"`
}

var persons = []Person{
	{Nama: "Risky", NIM: "123456", Alamat: "Jl. Mangga"},
	{Nama: "Akbar", NIM: "654321", Alamat: "Jl. Manggis"},
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {

	count := 0

	if r.Method == "POST" {
		GetNamePerson := Person{
			Nama: r.FormValue("Nama"),
		}

		for _, value := range persons {
			if strings.EqualFold(GetNamePerson.Nama, value.Nama) {
				response, err := json.Marshal(Person{
					Nama:   value.Nama,
					NIM:    value.NIM,
					Alamat: value.Alamat,
				})

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				count++

				w.Header().Set("Content-Type", "application/json")
				w.Write(response)
			}
		}

		if count == 0 {
			http.Error(w, "Nama Tidak ada", http.StatusMethodNotAllowed)
		}

	} else {
		http.Error(w, "Method Tidak tersedia", http.StatusMethodNotAllowed)
	}

}

func GetAllDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		personJson, _ := json.Marshal(persons)

		w.Header().Set("Content-Type", "application/json")
		w.Write(personJson)
	} else {
		http.Error(w, "Method Tidak tersedia", http.StatusMethodNotAllowed)
	}
}
