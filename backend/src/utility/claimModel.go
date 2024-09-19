package utility

import (
	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}
