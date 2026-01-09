package main

import (
	"log"
	"net/http"

	"github.com/bricefrisco/journalctl-gui/internal/httpapi"
)

func main() {
	if err := http.ListenAndServe(":8091", httpapi.NewRouter()); err != nil {
		log.Fatal(err)
	}
}
