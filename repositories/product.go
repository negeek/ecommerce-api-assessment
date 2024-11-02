package repositories

import (
	"context"
	"strconv"
	"time"

	"github.com/negeek/ecommerce-api-assessment/db"
	"github.com/negeek/ecommerce-api-assessment/utils"
)

type Product struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	StockQuantity int       `json:"stock_quantity"`
	DateCreated   time.Time `json:"date_created"`
	DateUpdated   time.Time `json:"date_updated"`
}

func (p *Product) Create() error {
	utils.Time(p, true)
	query := `
		INSERT INTO products (name, description, price, stock_quantity, date_created, date_updated) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`
	err := db.PostgreSQLDB.QueryRow(
		context.Background(), query, p.Name, p.Description, p.Price, p.StockQuantity, p.DateCreated, p.DateUpdated,
	).Scan(&p.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) FindByID(id int) error {
	query := `
		SELECT id, name, description, price, stock_quantity, date_created, date_updated 
		FROM products WHERE id = $1
	`
	row := db.PostgreSQLDB.QueryRow(context.Background(), query, id)
	return row.Scan(
		&p.ID, &p.Name, &p.Description, &p.Price, &p.StockQuantity, &p.DateCreated, &p.DateUpdated,
	)
}

func (p *Product) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := db.PostgreSQLDB.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) Put() error {
	utils.Time(p, false)
	query := `
		UPDATE products 
		SET name = $1, description = $2, price = $3, stock_quantity = $4, date_updated = $5 
		WHERE id = $6
	`
	_, err := db.PostgreSQLDB.Exec(
		context.Background(), query, p.Name, p.Description, p.Price, p.StockQuantity, p.DateUpdated, p.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) Patch() error {
	utils.Time(p, false)
	query := "UPDATE products SET "
	args := []interface{}{}
	argPosition := 1

	if p.Name != "" {
		query += "name = $" + strconv.Itoa(argPosition) + ", "
		args = append(args, p.Name)
		argPosition++
	}

	if p.Description != "" {
		query += "description = $" + strconv.Itoa(argPosition) + ", "
		args = append(args, p.Description)
		argPosition++
	}

	if p.Price != 0 {
		query += "price = $" + strconv.Itoa(argPosition) + ", "
		args = append(args, p.Price)
		argPosition++
	}

	if p.StockQuantity != 0 {
		query += "stock_quantity = $" + strconv.Itoa(argPosition) + ", "
		args = append(args, p.StockQuantity)
		argPosition++
	}

	query += "date_updated = $" + strconv.Itoa(argPosition)
	args = append(args, p.DateUpdated)
	argPosition++

	query += " WHERE id = $" + strconv.Itoa(argPosition)
	args = append(args, p.ID)

	// Execute the query
	_, err := db.PostgreSQLDB.Exec(context.Background(), query, args...)
	if err != nil {
		return err
	}
	return nil
}

func ProductList() ([]Product, error) {
	query := `SELECT id, name, description, price, stock_quantity, date_created, date_updated FROM products`
	rows, err := db.PostgreSQLDB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.StockQuantity, &product.DateCreated, &product.DateUpdated)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
