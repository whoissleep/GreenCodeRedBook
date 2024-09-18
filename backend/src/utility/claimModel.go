package utility

import (
	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	login    string `json:"login"`
	hashPass string `json:"pass"`
	jwt.StandardClaims
}
