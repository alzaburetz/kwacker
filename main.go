package main

import (
	"log"
	"net/http"

	"github.com/alzaburetz/kwacker/handlers"
)

func main() {
	http.HandleFunc("/user/add", handlers.PostUser)
	http.HandleFunc("/user/getall", handlers.GetUsers)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
