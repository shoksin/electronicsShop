package main

import (
	"ElectronicsStore/handlers"
	"ElectronicsStore/pkg/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	port := utils.GetEnv("PORT", "9999")

	r := mux.NewRouter()
	r.HandleFunc("/products", handlers.AllProducts)
	r.HandleFunc("/add", handlers.AddProduct)
	r.HandleFunc("/delete/{id}", handlers.DeleteProduct)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
