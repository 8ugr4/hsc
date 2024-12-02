package main

import (
	h "hsc/internal/http"
	"log"
)

// @todo: add pre-commit and golangci lint.
func main() {
	pages := []string{"https://www.google.com", "https://www.google.com"}
	status, err := h.Get(pages, "")
	if err != nil {
		log.Fatal("failed:", err)
	}
	for index := range status.Code {
		log.Println(status.Code[index], status.Text[index])
	}
}
