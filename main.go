package main

import (
	"log"
	"net/http"

	"github.com/alzaburetz/kwacker/handlers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	http.HandleFunc("/user/add", handlers.AddUser)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
