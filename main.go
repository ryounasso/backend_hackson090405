package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// Echo インスタンスをNew
	e := echo.New()

	// ルーティング設定
	e.GET("/books/:id", bookHandler)

	// サーバ起動
	e.Start(":" + os.Getenv("PORT"))
}

// ハンドラ
func bookHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, getBook())

}

func getBook() *Book {
	b := &Book{
		Title:  "Golangの本",
		Author: "Golangの本の著者",
	}
	return b
}
