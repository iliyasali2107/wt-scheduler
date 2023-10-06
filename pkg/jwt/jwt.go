package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "secret"

type CustomClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int, days int) (string, error) {
	claims := CustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * time.Duration(days))),
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func ParseToken(token string) (int, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) { return secretKey, nil })

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, fmt.Errorf("token is expired: %w", err)
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			return 0, fmt.Errorf("token is broken: %w", err)
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return 0, fmt.Errorf("invalid token signature %w", err)
		} else {
			return 0, fmt.Errorf("other error")
		}
	}

	claims, ok := parsedToken.Claims.(*CustomClaims)
	if !ok {
		return 0, fmt.Errorf("something wrong with claims")
	}

	return claims.UserId, nil
}
