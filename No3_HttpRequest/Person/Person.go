package Person

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Nama   string `json:"Nama"`
	NIM    string `json:"NIM"`
	Alamat string `json:"Alamat"`
}

var person []Person

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Person := Person{
			Nama:   r.FormValue("Nama"),
			NIM:    r.FormValue("NIM"),
			Alamat: r.FormValue("Alamat"),
		}

		person = append(person, Person)

		personJson, _ := json.Marshal(Person)

		w.Header().Set("Content-Type", "application/json")
		w.Write(personJson)
	} else {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
	}
}

func SemuaData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		personJson, _ := json.Marshal(person)

		w.Header().Set("Content-Type", "application/json")
        w.Write(personJson)
	} else {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
	}
}
