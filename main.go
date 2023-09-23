package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Calling root")
		io.WriteString(writer, "Hello world")
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
