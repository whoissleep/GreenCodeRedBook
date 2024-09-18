package main

import (
	"fmt"

	_ "github.com/GreenCodeBook/docs"
	"github.com/GreenCodeBook/src/middleware"
	"github.com/GreenCodeBook/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	appGroup := app.Group("/moder")
	appGroup.Use(middleware.Aunthorization)
	db := storage.Init()
	fmt.Println(db)
}
