package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	server := http.ListenAndServe(":3000", router)
	log.Fatal(server)
	fmt.Printf("El servidor esta corriendo...")
}
