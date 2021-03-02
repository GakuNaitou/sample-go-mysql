package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SampleUser struct {
	gorm.Model
	Name string
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// t := &Template{
	// 	templates: template.Must(template.ParseGlob("views/*.html")),
	// }
	e := echo.New()
	// e.Renderer = t

	// dsn := fmt.Sprintf(
	// 	"%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	os.Getenv("MYSQL_USER"),
	// 	os.Getenv("MYSQL_PASSWORD"),
	// 	os.Getenv("MYSQL_DB"),
	// )
	// sqlDB, err := sql.Open(os.Getenv("MYSQL_HOST"), dsn)
	// if err != nil {
	// 	panic(err)
	// }
	// gormDB, err := gorm.Open(mysql.New(mysql.Config{
	// 	Conn: sqlDB,
	// }), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }

	// gormDB.AutoMigrate(&SampleUser{})

	e.GET("/", func(c echo.Context) error {
		// var sampleUsers []SampleUser
		// gormDB.Find(&sampleUsers)

		// ViewData := struct {
		// 	SampleUsers []SampleUser
		// }{
		// 	SampleUsers: sampleUsers,
		// }

		return c.Render(http.StatusOK, "main.html", "hoge")
	})

	e.POST("/", func(c echo.Context) error {
		// sampleUserName := c.FormValue("name")
		// gormDB.Create(&SampleUser{Name: sampleUserName})

		return c.Redirect(http.StatusFound, "/")
	})

	e.Logger.Fatal(e.Start(":9000"))
}
