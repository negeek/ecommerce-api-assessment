package repositories

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/negeek/ecommerce-api-assessment/db"
	"github.com/negeek/ecommerce-api-assessment/utils"
)

type Order struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Products    []Product `json:"products"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

func (o *Order) Create() error {
	utils.Time(o, true)
	productsJSON, err := json.Marshal(o.Products)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO orders (user_id, products, total_amount, status, date_created, date_updated)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`
	err = db.PostgreSQLDB.QueryRow(
		context.Background(), query, o.UserID, productsJSON, o.TotalAmount, o.Status, time.Now(), time.Now(),
	).Scan(&o.ID)
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) FindByID(id int) error {
	query := `
		SELECT id, user_id, products, total_amount, status, date_created, date_updated 
		FROM orders WHERE id = $1
	`
	row := db.PostgreSQLDB.QueryRow(context.Background(), query, id)

	var productsJSON []byte
	err := row.Scan(
		&o.ID, &o.UserID, &productsJSON, &o.TotalAmount, &o.Status, &o.DateCreated, &o.DateUpdated,
	)
	if err != nil {
		return err
	}

	err = json.Unmarshal(productsJSON, &o.Products)
	return err
}

func (o *Order) Put() error {
	utils.Time(o, false)
	productsJSON, err := json.Marshal(o.Products)
	if err != nil {
		return err
	}

	query := `
		UPDATE orders 
		SET user_id = $1, products = $2, total_amount = $3, status = $4, date_updated = $5 
		WHERE id = $6
	`
	_, err = db.PostgreSQLDB.Exec(
		context.Background(), query, o.UserID, productsJSON, o.TotalAmount, o.Status, o.DateUpdated, o.ID,
	)
	return err
}

func (o *Order) Delete(id int) error {
	query := "DELETE FROM orders WHERE id = $1"
	_, err := db.PostgreSQLDB.Exec(context.Background(), query, id)
	return err
}

func (o *Order) Patch() error {
	utils.Time(o, false)
	query := "UPDATE orders SET "
	args := []interface{}{}
	argPosition := 1

	// Check and add fields to update
	if o.UserID != 0 {
		query += "user_id = $" + strconv.Itoa(argPosition) + ", "
		args = append(args, o.UserID)
		argPosition++
	}

	if o.TotalAmount != 0 {
		query += "total_amount = $" + strconv.Itoa(argPosition) + ", "
		args = append(args, o.TotalAmount)
		argPosition++
	}

	if o.Status != "" {
		query += "status = $" + strconv.Itoa(argPosition) + ", "
		args = append(args, o.Status)
		argPosition++
	}

	if len(o.Products) > 0 {
		productsJSON, err := json.Marshal(o.Products)
		if err != nil {
			return err
		}
		query += "products = $" + strconv.Itoa(argPosition) + ", "
		args = append(args, productsJSON)
		argPosition++
	}

	query += "date_updated = $" + strconv.Itoa(argPosition)
	args = append(args, o.DateUpdated)
	argPosition++

	query += " WHERE id = $" + strconv.Itoa(argPosition)
	args = append(args, o.ID)

	// Execute the query
	_, err := db.PostgreSQLDB.Exec(context.Background(), query, args...)
	return err
}

func FindOrdersByUserID(userID int) ([]Order, error) {
	query := `
		SELECT id, user_id, products, total_amount, status, date_created, date_updated
		FROM orders WHERE user_id = $1
	`
	rows, err := db.PostgreSQLDB.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed after use

	var orders []Order

	for rows.Next() {
		var productsJSON []byte
		order := Order{}
		err := rows.Scan(
			&order.ID, &order.UserID, &productsJSON, &order.TotalAmount, &order.Status,
			&order.DateCreated, &order.DateUpdated,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(productsJSON, &order.Products)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
