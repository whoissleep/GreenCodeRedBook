package utility

import (
	"github.com/dgrijalva/jwt-go"
)

func ParseToken(tokenString string) (claims *Claim, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("it_is_my_password"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claim)

	if !ok {
		return nil, err
	}

	return claims, nil
}