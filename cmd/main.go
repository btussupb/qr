package main

import (
	"log"
	"net/http"

	wayToTrasport "github.com/btussupb/qr/internal/transport"
)

func main() {
	log.Print("http://localhost:8111")

	if err := http.ListenAndServe(":8111", wayToTrasport.Router()); err != nil {
		log.Fatal(err)
	}
}
