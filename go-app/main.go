package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
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

// func main() {
// 	t := &Template{
// 		templates: template.Must(template.ParseGlob("views/*.html")),
// 	}
// 	e := echo.New()
// 	e.Renderer = t

// 	dsn := fmt.Sprintf(
// 		"%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local",
// 		os.Getenv("MYSQL_USER"),
// 		os.Getenv("MYSQL_PASSWORD"),
// 		os.Getenv("MYSQL_DB"),
// 	)
// 	sqlDB, err := sql.Open(os.Getenv("MYSQL_HOST"), dsn)
// 	if err != nil {
// 		panic(err)
// 	}
// 	gormDB, err := gorm.Open(mysql.New(mysql.Config{
// 		Conn: sqlDB,
// 	}), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	gormDB.AutoMigrate(&SampleUser{})

// 	e.GET("/", func(c echo.Context) error {
// 		var sampleUsers []SampleUser
// 		gormDB.Find(&sampleUsers)

// 		ViewData := struct {
// 			SampleUsers []SampleUser
// 		}{
// 			SampleUsers: sampleUsers,
// 		}

// 		return c.Render(http.StatusOK, "main.html", ViewData)
// 	})

// 	e.POST("/", func(c echo.Context) error {
// 		sampleUserName := c.FormValue("name")
// 		gormDB.Create(&SampleUser{Name: sampleUserName})

// 		return c.Redirect(http.StatusFound, "/")
// 	})

// 	e.Logger.Fatal(e.Start(":9000"))
// }

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

		// return c.Render(http.StatusOK, "main.html", "hoge")
		return c.String(http.StatusOK, "Hello, World!")
	})

	// e.POST("/", func(c echo.Context) error {
	// 	sampleUserName := c.FormValue("name")
	// 	gormDB.Create(&SampleUser{Name: sampleUserName})

	// 	return c.Redirect(http.StatusFound, "/")
	// })

	e.Logger.Fatal(e.Start(":9000"))
}

// func main() {
// 	log.Print("starting server...")
// 	http.HandleFunc("/", handler)

// 	// Determine port for HTTP service.
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 		log.Printf("defaulting to port %s", port)
// 	}

// 	// Start HTTP server.
// 	log.Printf("listening on port %s", port)
// 	if err := http.ListenAndServe(":"+port, nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	name := os.Getenv("NAME")
// 	if name == "" {
// 		name = "World"
// 	}
// 	fmt.Fprintf(w, "Hello %s!\n", name)
// }
