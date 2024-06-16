package handlers

import (
	"ElectronicsStore/db"
	"ElectronicsStore/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var data *db.DB = db.NewDatabase()

func AllProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := data.Query("select * from Products")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	products := []models.Product{}
	for rows.Next() {
		p := models.Product{}
		err := rows.Scan(&p.ID, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	tmpl, _ := template.ParseFiles("static/template/index.html")
	err = tmpl.Execute(w, products)
	if err != nil {
		log.Print(err.Error())
	}
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		model := r.FormValue("model")
		company := r.FormValue("company")
		price, _ := strconv.Atoi(r.FormValue("price"))

		p := models.Product{Model: model, Company: company, Price: price}

		err := data.CreateProduct(&p)
		if err != nil {
			log.Print(err.Error())
		}
		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}

	tmpl, err := template.ParseFiles("static/template/create.html")
	if err != nil {
		http.Error(w, "Ошибка при загрузке шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Не передан ID продукта", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Ошибка обработки ID для удаления"+err.Error(), http.StatusInternalServerError)
	}

	err = data.DeleteProduct(id)
	if err != nil {
		http.Error(w, "Ошибка удаления: "+err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)

}
