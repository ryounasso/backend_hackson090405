package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

type Book struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		rows, err := db.Query(`SELECT id, name FROM mybook`)
		var books []Book
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var id int64
			var name string
			err = rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			book := Book{Id: id, Name: name}
			books = append(books, book)
		}

		fmt.Println(&rows)
		return c.JSON(http.StatusOK, books)
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
