package login

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	Service struct{}

	Claims struct {
		CustomerID string `json:"customer_id"`
		jwt.StandardClaims
	}
)

func NewService() *Service {
	return &Service{}
}

func (svc *Service) Login(_ context.Context, customerID string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	expirationHours, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_HOURS"))
	if err != nil {
		return "", err
	}

	claims := Claims{
		CustomerID: customerID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expirationHours) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
