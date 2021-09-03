package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	DB "github.com/ryounasso/backend_hackson090405/db"
)

func main() {
	db := DB.GetDBConnection()
	DB.DBMigrate(db)

	var todos []DB.Todo

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		db.Find(&todos)
		return c.JSON(http.StatusOK, todos)
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}
	e.Logger.Fatal(e.Start(":" + PORT))
}
