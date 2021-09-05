package handler

import (
	"net/http"

	"github.com/labstack/echo"
	DB "github.com/ryounasso/backend_hackson090405/db"
)

func GetCompletedTodoNum(c echo.Context) error {
	var todos []DB.Todo
	db := DB.GetDBConnection()
	userId := c.Param("userId")

	// db.Where("user_id = ?", userId).Find(&todos)
	db.Where(&DB.Todo{UserId: userId, IsCompleted: true}).Find(&todos)

	return c.JSON(http.StatusOK, len(todos))
}
