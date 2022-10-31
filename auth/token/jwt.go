package token

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTTokenGen struct {
	PrivateKey *rsa.PrivateKey
	Issuer     string
	NowFunc    func() time.Time
}

func (t *JWTTokenGen) GenerateToken(accountID string, expire time.Duration) (string, error) {
	nowSec := t.NowFunc()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.RegisteredClaims{
		Issuer:    t.Issuer,
		IssuedAt:  jwt.NewNumericDate(nowSec),
		ExpiresAt: jwt.NewNumericDate(nowSec.Add(expire)),
		Subject:   accountID,
	})

	return tkn.SignedString(t.PrivateKey)
}
