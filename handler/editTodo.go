package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	DB "github.com/ryounasso/backend_hackson090405/db"
)

func EditTodo(c echo.Context) error {
	var todo DB.Todo
	db := DB.GetDBConnection()

	todoId := c.FormValue("todoId")
	title := c.FormValue("title")
	description := c.FormValue("description")
	fmt.Println(todoId, title, description)

	db.Model(&DB.Todo{}).Where("id = ?", todoId).Updates(DB.Todo{Title: title, Description: description})
	return c.JSON(http.StatusOK, todo)
}
