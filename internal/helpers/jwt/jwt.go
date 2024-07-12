package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/danzBraham/eniqilo-store/internal/errors/autherror"
	"github.com/golang-jwt/jwt/v5"
)

var key = []byte(os.Getenv("JWT_SECRET"))

type CustomClaims struct {
	UserID string
	jwt.RegisteredClaims
}

func GenerateToken(ttl time.Duration, userID string) (string, error) {
	now := time.Now()
	expiry := now.Add(ttl)

	claims := &CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiry),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

type JWTPayload struct {
	UserID string
}

func VerifyToken(tokenString string) (*JWTPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Method.Alg())
		}
		return key, nil
	})
	if err != nil {
		return nil, autherror.ErrInvalidToken
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, autherror.ErrUnknownClaims
	}
	return &JWTPayload{
		UserID: claims.UserID,
	}, nil
}
