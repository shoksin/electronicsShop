package db

import (
	"ElectronicsStore/models"
	"database/sql"
	"fmt"
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

func (db *DB) CreateProduct(p *models.Product) error {
	_, err := db.Query(fmt.Sprintf("INSERT INTO products(model, company, price) values ('%s', '%s', '%d')", p.Model, p.Company, p.Price))
	if err != nil {
		panic(err)
	}
	return err
}
