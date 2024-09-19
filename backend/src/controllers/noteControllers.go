package controllers

import (
	"github.com/GreenCodeBook/src/models"
	"github.com/GreenCodeBook/src/services"
	"github.com/GreenCodeBook/src/utility"
	"github.com/gofiber/fiber/v2"
)

// AddUser Adding User by json
//
//	@Summary		Add Note by json
//	@Description	Add Note by jso in detail
//	@Tags			Notes
//	@Accept			json
//	@Produce		json
//	@Param			request			body		models.NoteDto	true	"Request of Creating Note Object"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/note [post]
func AddNote(app *fiber.App, noteService services.NoteService) fiber.Router {
	return app.Post("/note", func(c *fiber.Ctx) error {
		var noteDto models.NoteDto
		err := c.BodyParser(&noteDto)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("User format is not valid")
		}
		note, err := noteDto.MapToNote()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}
		err = noteService.AddNote(note)
		if err == nil {
			return c.Status(fiber.StatusOK).JSON("added")
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}
	})
}

// Get Users Notes
//
//	@Summary		Add Note by json
//	@Description	Add Note by jso in detail
//	@Tags			Notes
//	@Accept			json
//	@Produce		json
//	@Param			token		header		string	true	"token"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/noteGetAll [get]
func GetNotesByToken(app *fiber.App, noteService services.NoteService) fiber.Router {
	return app.Get("/noteGetAll", func(c *fiber.Ctx) error {
		token := c.Get("token")

		claims, er := utility.ParseToken(token)
		if er != nil {
			return c.Status(fiber.StatusBadRequest).JSON(er)
		}

		id := claims.ID

		ans, er := noteService.GetAllNotes(id)
		if er != nil {
			return c.Status(fiber.StatusNotFound).JSON(string(er.Error()))
		}

		return c.Status(fiber.StatusOK).JSON(ans)
	})
}
