package db

import (
	"ElectronicsStore/models"
	"database/sql"
	"fmt"
	"net/http"
)

type DB struct {
	*sql.DB
}

func NewDatabase() *DB {
	connStr := "user=postgres password=sinoflan06 dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	return &DB{db}
}

func (db *DB) GetAllProducts(w http.ResponseWriter) ([]models.Product, error) {
	rows, err := db.Query("select * from Products")
	if err != nil {
		http.Error(w, "internal error with database ", http.StatusInternalServerError)
		return nil, err
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

	return products, nil
}

func (db *DB) CreateProduct(p *models.Product) error {
	_, err := db.Query(fmt.Sprintf("INSERT INTO products(model, company, price) values ('%s', '%s', '%d')", p.Model, p.Company, p.Price))
	if err != nil {
		panic(err)
	}
	return err
}

func (db *DB) DeleteProduct(id int) error {
	_, err := db.Query(fmt.Sprintf("DELETE FROM products WHERE id = '%d'", id))
	return err
}
