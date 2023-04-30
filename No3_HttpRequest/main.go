package main

import (
    "net/http"
    "Person"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/nama", Person.PersonHandler)
    mux.HandleFunc("/semuadata", Person.GetAllDataHandler)

    http.ListenAndServe(":5050", mux)
}