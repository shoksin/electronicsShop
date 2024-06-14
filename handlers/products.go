package handlers

import (
	"ElectronicsStore/db"
	"ElectronicsStore/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

		fmt.Println(model, company, price)

		p := models.Product{Model: model, Company: company, Price: price}

		err := data.CreateProduct(&p)
		if err != nil {
			log.Print(err.Error())
		}
		var a = 5
		fmt.Println(a)
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

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
