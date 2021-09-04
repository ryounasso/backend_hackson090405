package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	DB "github.com/ryounasso/backend_hackson090405/db"
)

func AddTodo(c echo.Context) error {
	db := DB.GetDBConnection()
	title := c.FormValue("title")
	description := c.FormValue("description")
	fmt.Println(title, description)
	todo := DB.Todo{}

	todo.ID = 10
	todo.Title = "aaaa"
	todo.Description = "daaa"
	todo.IsCompleted = false

	db.Create(&todo)
	return c.JSON(http.StatusOK, todo)
}
