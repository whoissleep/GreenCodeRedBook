package main

import (
	_ "github.com/GreenCodeBook/docs"
	"github.com/GreenCodeBook/src/controllers"
	"github.com/GreenCodeBook/src/middleware"
	"github.com/GreenCodeBook/src/services"
	"github.com/GreenCodeBook/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

// @title			2lab API
// @version		1.0
// @description	This is a sample swagger for GreenHack
// @termsOfService	http://swagger.io/terms/
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	appGroup := app.Group("/moder")
	appGroup.Use(middleware.Aunthorization)
	db := storage.Init()
	//defer storage.Close()
	userService := services.NewUserSevice(db)
	controllers.AddUser(app, userService)
	controllers.LoginUser(app, userService)
	app.Listen(":3001")
}

/*
{
  "email": "y@ya.ru",
  "name": "Василий Степанович",
  "pass": "123",
  "phone": "8999999999"
}


{
  "email": "tom@ya.ru",
  "name": "Tom Rom",
  "pass": "123",
  "phone": "777777777"
}

*/
