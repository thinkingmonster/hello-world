package main

import (
	"fmt"
	"myapp/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("Strating server on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
