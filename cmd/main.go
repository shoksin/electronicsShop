package main

import (
	"ElectronicsStore/handlers"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/products", handlers.AllProducts)
	http.HandleFunc("/add", handlers.AddProduct)
	http.ListenAndServe(":8080", nil)
}
