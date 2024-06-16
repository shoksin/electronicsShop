package main

import (
	"ElectronicsStore/handlers"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", handlers.AllProducts)
	r.HandleFunc("/add", handlers.AddProduct)
	r.HandleFunc("/delete/{id}", handlers.DeleteProduct)
	// Обслуживание статических файлов
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":8080", r)
}
