package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/lib/pq"
)

type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type Person struct {
	Model
	Name string
	Age  int
}

func main() {
	db := GetDBConnection()

	db.AutoMigrate(&Person{})
	var persons []Person

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		db.Find(&persons)
		return c.JSON(http.StatusOK, persons)
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}
	e.Logger.Fatal(e.Start(":" + PORT))
}

func GetDBConnection() *gorm.DB {
	url := os.Getenv("DATABASE_URL")
	connection, err := pq.ParseURL(url)
	if err != nil {
		log.Fatal(err)
	}
	connection += " sslmode=disable"
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
