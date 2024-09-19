package services

import (
	"errors"
	"fmt"
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
	fmt.Println("here2")
	erdb := r.db.Create(&user)
	if erdb.Error != nil {
		fmt.Println(erdb)
		return erdb.Error
	}
	return nil
}

func (r userService) Login(user models.User) (string, error) {
	var dbUser models.User
	erdb := r.db.Where("email = ?", user.Email).First(&dbUser)
	fmt.Println("из бд:", dbUser)
	if erdb.Error != nil {
		return "", erdb.Error
	}
	//fmt.Println("here in login before pass")
	if !utility.Compare(user.HashPas, dbUser.HashPas) {
		//fmt.Print("here in login bad pass")
		return "", errors.New("invalid password")
	}

	exTime := time.Now().Add(5 * time.Minute)
	fmt.Println("here in Login token 1", dbUser.Id, dbUser.Role)
	claims := &utility.Claim{
		ID:   int(dbUser.Id),
		Role: dbUser.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   dbUser.Email,
			ExpiresAt: exTime.Unix(),
		},
	}
	fmt.Println("here in Login token 2")
	fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var SecretKey = os.Getenv("SECRETKEY")
	fmt.Println("here in Login token 3", SecretKey, " <-  ->", os.Getenv("SECRETKEY"))
	tokenString, err := token.SignedString([]byte(SecretKey))

	claims, er := utility.ParseToken(tokenString)
	if er != nil {
		fmt.Println("error in parse ", er)
	}
	fmt.Println("after parse in login", claims)

	fmt.Println("here in Login error", err)
	if err != nil {
		return "", errors.New("token couldnt be generated")
	}
	return tokenString, nil
}
func NewUserSevice(db *gorm.DB) UserService {
	return &userService{db: db}
}
