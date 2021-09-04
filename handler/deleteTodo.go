package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	DB "github.com/ryounasso/backend_hackson090405/db"
)

func DeleteTodo(c echo.Context) error {
	var todo DB.Todo
	db := DB.GetDBConnection()

	todoId, _ := strconv.Atoi(c.Param("todoId"))
	todo.ID = uint(todoId)
	db.Delete(&todo)
	return c.JSON(http.StatusOK, "delete ok")
}
