package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	ServerID := os.Getenv("SERVERID")
	io.WriteString(w, "msg from server"+ServerID)
}

func main() {
	http.HandleFunc("/getResult", Handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
