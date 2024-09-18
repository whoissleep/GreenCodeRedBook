package utility

import (
	"golang.org/x/crypto/bcrypt"
)

func GetHesh(pass string) (string, error) {
	hash, er := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(hash), er
}

func Compare(pass string, hash string) bool {
	er := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return er == nil

}
