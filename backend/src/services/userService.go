package services

import (
	"errors"
	"os"
	"time"

	"github.com/GreenCodeBook/src/models"
	"github.com/GreenCodeBook/src/utility"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type userService struct {
	db *gorm.DB
}

type UserService interface {
	AddUser(user models.User) error
	Login(user models.User) (string, error)
}

func (r userService) AddUser(user models.User) error {
	user.HashPas, _ = utility.GetHesh(user.HashPas)
	erdb := r.db.Create(&user)
	if erdb.Error != nil {
		return erdb.Error
	}
	return nil
}

func (r userService) Login(user models.User) (string, error) {
	var dbUser models.User
	erdb := r.db.Where("email = ?", user.Email).First(&dbUser)
	if erdb.Error != nil {
		return "", erdb.Error
	}
	if !utility.Compare(user.HashPas, dbUser.HashPas) {
		return "", errors.New("invalid password")
	}

	exTime := time.Now().Add(5 * time.Minute)
	claims := &utility.Claim{
		ID:   int(dbUser.Id),
		Role: dbUser.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   dbUser.Email,
			ExpiresAt: exTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var SecretKey = os.Getenv("SECRETKEY")
	tokenString, err := token.SignedString([]byte(SecretKey))

	claims, er := utility.ParseToken(tokenString)
	if er != nil {
		panic(er)
	}

	if err != nil {
		return "", errors.New("token couldnt be generated")
	}
	return tokenString, nil
}
func NewUserSevice(db *gorm.DB) UserService {
	return &userService{db: db}
}
