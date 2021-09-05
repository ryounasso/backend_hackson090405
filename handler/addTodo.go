package handler

import (
	"net/http"

	"github.com/labstack/echo"
	DB "github.com/ryounasso/backend_hackson090405/db"
)

func AddTodo(c echo.Context) error {
	db := DB.GetDBConnection()
	userId := c.FormValue("userId")
	title := c.FormValue("title")
	description := c.FormValue("description")
	todo := DB.Todo{}

	todo.UserId = userId
	todo.Title = title
	todo.Description = description
	todo.IsCompleted = false

	db.Create(&todo)
	return c.JSON(http.StatusOK, &todo)
}
