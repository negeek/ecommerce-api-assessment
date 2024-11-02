package services

import (
	"errors"

	userenum "github.com/negeek/ecommerce-api-assessment/enums"
	"github.com/negeek/ecommerce-api-assessment/repositories"
	"github.com/negeek/ecommerce-api-assessment/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (u *UserService) Register(user *repositories.User) error {
	refuser := &repositories.User{}
	err := refuser.FindByEmail(user.Email)
	if err == nil {
		return errors.New("email already exist")
	}
	roleCheck := &userenum.UserRole{}
	if !roleCheck.IsValid(user.Role) {
		return errors.New("invalid user role")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return user.Create()
}

func (u *UserService) GenerateJWT(user *repositories.User) (string, error) {
	return utils.CreateJwtToken(user.ID, user.Email)
}

func (u *UserService) Login(user *repositories.User) (string, error) {
	password := user.Password
	err := user.FindByEmail(user.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := u.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
