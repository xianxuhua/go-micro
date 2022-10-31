package token

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

func (v *JWTTokenVerifier) Verify(token string) (string, error) {
	tkn, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
		return v.PublicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("can not parse token: %v", err)
	}
	if !tkn.Valid {
		return "", fmt.Errorf("token not valid")
	}
	clm, ok := tkn.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return "", fmt.Errorf("token claim is not StandardClaims")
	}

	if err := clm.Valid(); err != nil {
		return "", fmt.Errorf("claim is not Valid: %v", err)
	}

	return clm.Subject, nil
}
