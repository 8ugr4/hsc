package main

import (
	h "hsc/internal/http"
	"log"
)

func main() {
	page := "https://www.google.com"
	auth := ""

	status, err := h.Get(page, auth)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(status.Code, status.Text)
}
