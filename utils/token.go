package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	jwt.RegisteredClaims
	ID    int
	Email string
}

func CreateJwtToken(id int, email string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour) // Set to expire in 24 hours

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
		ID:    id,
		Email: email,
	})

	// Create the actual JWT token
	signedString, err := token.SignedString([]byte(os.Getenv("AUTH_KEY")))

	if err != nil {
		return "", err
	}

	return signedString, nil
}

func VerifyJwt(jwtToken string) (*UserClaim, error) {
	// Parse and validate the JWT token
	token, err := jwt.ParseWithClaims(jwtToken, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("AUTH_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*UserClaim), nil
}
