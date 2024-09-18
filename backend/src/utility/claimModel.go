package utility

import (
	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Login    string `json:"login"`
	HashPass string `json:"pass"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
