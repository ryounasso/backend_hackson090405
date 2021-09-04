package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/ryounasso/backend_hackson090405/handler"
)

// func dumpHandler(c echo.Context, reqBody, resBody []byte) {
// 	fmt.Fprintf(os.Stdout, "Request: %+v\n", string(reqBody))
// }

func main() {
	e := echo.New()
	// e.HideBanner = true
	// e.HidePort = true

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

	e.GET("/todos/:userId", handler.GetTodos)
	e.POST("/todos/edit", handler.EditTodo)
	e.POST("/todos/add", handler.AddTodo)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}
	e.Logger.Fatal(e.Start(":" + PORT))
}
