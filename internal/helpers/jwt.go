package helpers

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func CreateJWT(staffId string) (string, error) {
	jwtExp, _ := strconv.Atoi(os.Getenv("JWT_EXP"))
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	claims := &CustomClaims{
		staffId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtExp) * time.Hour)),
			Issuer:    "eniqilo-store-server",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
