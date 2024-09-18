package controllers

// AddUser Adding User by json
//
//	@Summary		Add User by json
//	@Description	Add User by jso in detail
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request			body		dtos.UserDto	true	"Request of Creating User Object"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/user [post]
