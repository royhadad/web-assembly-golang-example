package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 9000...")
	log.Fatalln(http.ListenAndServe(":9000", http.FileServer(http.Dir("./assets"))))
}
