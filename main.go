package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
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
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	return db
}
