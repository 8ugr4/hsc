package main

import (
	h "hsc/internal/http"
	"log"
)

// @todo: add pre-commit and golangci lint.
func main() {
	pages := []string{"https://www.google.com", "https://www.google.com"}
	err := h.Get(pages, "")
	if err != nil {
		log.Fatal(err)
	}
}
