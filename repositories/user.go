package repositories

import (
	"context"
	"time"

	"github.com/negeek/ecommerce-api-assessment/db"
	userenum "github.com/negeek/ecommerce-api-assessment/enums"
	"github.com/negeek/ecommerce-api-assessment/utils"
)

type User struct {
	ID          int       `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Role        string    `json:"role"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

func (u *User) IsAdmin() bool {
	return u.Role == userenum.Admin
}

func (u *User) Create() error {
	utils.Time(u, true)
	query := `
		INSERT INTO users (email, password, role, date_created, date_updated)
		VALUES ($1, $2, $3, $4, $5) RETURNING id
	`
	err := db.PostgreSQLDB.QueryRow(
		context.Background(), query, u.Email, u.Password, u.Role, u.DateCreated, u.DateUpdated,
	).Scan(&u.ID)

	if err != nil {
		return err
	}
	return nil
}

func (u *User) FindByID(id int) error {
	query := "SELECT id, password, email, role, date_created, date_updated FROM users WHERE id = $1"
	row := db.PostgreSQLDB.QueryRow(context.Background(), query, id)

	err := row.Scan(&u.ID, &u.Password, &u.Email, &u.Role, &u.DateCreated, &u.DateUpdated)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) FindByEmail(email string) error {
	query := "SELECT id, password, email, role, date_created, date_updated FROM users WHERE email = $1"
	row := db.PostgreSQLDB.QueryRow(context.Background(), query, email)

	err := row.Scan(&u.ID, &u.Password, &u.Email, &u.Role, &u.DateCreated, &u.DateUpdated)
	if err != nil {
		return err
	}
	return nil
}
