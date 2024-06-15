package main

import (
	"ElectronicsStore/handlers"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// 	http.HandleFunc("/products", handlers.AllProducts)
	// 	http.HandleFunc("/add", handlers.AddProduct)
	// 	http.HandleFunc("/delete/{id}", handlers.DeleteProduct)
	// 	http.ListenAndServe(":8080", nil)
	// }

	r := mux.NewRouter()
	r.HandleFunc("/products", handlers.AllProducts)
	r.HandleFunc("/add", handlers.AddProduct)
	r.HandleFunc("/delete/{id}", handlers.DeleteProduct)
	http.ListenAndServe(":8080", r)
}
