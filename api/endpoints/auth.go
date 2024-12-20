package endpoints

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserUuid string `json:"user_uuid"`
	jwt.RegisteredClaims
}

func (e *Endpoint) createJwt(userUuid string) (string, error) {
	// Create claims with multiple fields populated
	claims := Claims{
		UserUuid: userUuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    e.config.App.Name,
			Subject:   "authentication",
		},
	}

	// Create a new JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(e.config.Auth.SigningKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (e *Endpoint) parseJwt(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return []byte(e.config.Auth.SigningKey), nil
	})
 }