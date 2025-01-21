package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func getResultHandler(w http.ResponseWriter, r *http.Request) {
	ServerID := os.Getenv("SERVERID")
	io.WriteString(w, "msg from server"+ServerID)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	ServerID := os.Getenv("SERVERID")
	io.WriteString(w, "Welcome to server "+ServerID)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/getResult", getResultHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
