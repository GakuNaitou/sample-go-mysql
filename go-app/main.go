package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SampleUser struct {
	gorm.Model
	Name string
}

func main() {
	e := echo.New()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DB"),
	)

	sqlDB, err := sql.Open(os.Getenv("MYSQL_HOST"), dsn)
	if err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	gormDB.AutoMigrate(&SampleUser{})
	gormDB.Create(&SampleUser{Name: "sample name"})

	e.GET("/", func(c echo.Context) error {
		var sampleUser SampleUser
		gormDB.First(&sampleUser)

		return c.String(http.StatusOK,
			fmt.Sprintf("%s", sampleUser.Name),
		)
	})

	e.Logger.Fatal(e.Start(":9000"))
}
