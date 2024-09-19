package controllers

import (
	"github.com/GreenCodeBook/src/models"
	"github.com/GreenCodeBook/src/services"
	"github.com/gofiber/fiber/v2"
)

// AddUser Adding User by json
//
//	@Summary		Add User by json
//	@Description	Add User by jso in detail
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request			body		models.UserDto	true	"Request of Creating User Object"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/user [post]
func AddUser(app *fiber.App, userService services.UserService) fiber.Router {
	return app.Post("/user", func(c *fiber.Ctx) error {
		var UserDto models.UserDto
		var User models.User
		err := c.BodyParser(&UserDto)
		User = UserDto.MapNewUserDtoToUser()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("User format is not valid")
		}
		err = userService.AddUser(User)
		if err == nil {
			return c.Status(fiber.StatusOK).JSON("added")
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}
	})
}

// Login User Adding User by json
//
//		@Summary		Login User by json
//		@Description	Login User by jso in detail
//		@Tags			Users
//		@Accept			json
//		@Produce		json
//		@Param			login			header		string	true	"Login"
//	 	@Param			password		header		string	true	"Password"
//		@Failure		400				{string}	string
//		@Failure		404				{string}	string
//		@Success		200				{string}	string
//		@Router			/login [get]
func LoginUser(app *fiber.App, userService services.UserService) fiber.Router {
	return app.Get("/login", func(c *fiber.Ctx) error {
		var userDto models.LoginUserDto
		userDto.Email = c.Get("login")
		userDto.Password = c.Get("password")
		var user models.User = userDto.MapToUser()
		token, err := userService.Login(user)
		if err == nil {
			return c.Status(fiber.StatusOK).JSON(token)
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
	})
}
