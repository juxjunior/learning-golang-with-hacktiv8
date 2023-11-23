package main

import (
	"fmt"
	"net/http"
)

// Contoh web base app

var PORT = ":8080" // value tidak boleh kosong

func main() {
	mainUntukWeb()
}

func mainUntukWeb() {
	http.HandleFunc("/", greet)

	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello Web Client..."
	fmt.Fprint(w, msg)
}

// jalankan di browser http://localhost:8080/ ketika file ini dieksekusi
