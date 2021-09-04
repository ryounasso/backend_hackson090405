package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	DB "github.com/ryounasso/backend_hackson090405/db"
	"github.com/ryounasso/backend_hackson090405/handler"
)

// func dumpHandler(c echo.Context, reqBody, resBody []byte) {
// 	fmt.Fprintf(os.Stdout, "Request: %+v\n", string(reqBody))
// }

func main() {
	// e.HideBanner = true
	// e.HidePort = true
	db := DB.GetDBConnection()
	DB.DBMigrate(db)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	// e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
	// 	Skipper: func(c echo.Context) bool {
	// 		if c.Request().Header.Get("X-Debug") == "" {
	// 			return true
	// 		}
	// 		return false
	// 	},
	// 	Handler: dumpHandler,
	// }))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hackson090405_不退転")
	})
	e.GET("/todos/:userId", handler.GetTodos)
	e.POST("/todos/add", handler.AddTodo)
	e.POST("/todos/edit/:todoId", handler.EditTodo)
	e.POST("/todos/toggle_todo/:todoId", handler.ToggleTodo)
	e.DELETE("/todos/delete/:todoId", handler.DeleteTodo)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}
	e.Logger.Fatal(e.Start(":" + PORT))
}
