package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	DB "github.com/ryounasso/backend_hackson090405/db"
)

func EditTodo(c echo.Context) error {
	var todo DB.Todo
	db := DB.GetDBConnection()

	todoId, _ := strconv.Atoi(c.Param("todoId"))
	title := c.FormValue("title")
	description := c.FormValue("description")
	todo.ID = uint(todoId)
	todo.Title = title
	todo.Description = description
	db.Model(&todo).Updates(DB.Todo{Title: title, Description: description})

	return c.JSON(http.StatusOK, todo)
}
