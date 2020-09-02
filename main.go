// при вставке текста updatePost() - проблемы с кавычками " ' "

package main

import (
	"database/sql"
	"log"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server - структура сервера
type Server struct {
	db *sql.DB
}

// Post - структура одного поста
type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Link    string `json:"link"`
	Comment string `json:"comment"`
}

var (
	tmplList   = template.Must(template.New("MyTemplate").ParseFiles("./templates/list.html"))
	tmplSingle = template.Must(template.New("MyTemplate").ParseFiles("./templates/single.html"))
	tmplEdit   = template.Must(template.New("MyTemplate").ParseFiles("./templates/edit.html"))
	tmplCreate = template.Must(template.New("MyTemplate").ParseFiles("./templates/create.html"))
)

func main() {
	DSN := "root:1234@tcp(localhost:3306)/blog_posts?charset=utf8"
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	s := Server{
		db: db,
	}

	e := echo.New()
	e.HideBanner = true
	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.handlePostsList)
	e.GET("/post/:id", s.handleSinglePostGET)
	e.POST("/post/:id", s.handleSinglePostPOST)
	e.GET("/edit/:id", s.handleEditPostGET)
	e.GET("/create/new", s.handleCreatePostGET)
	e.POST("/create/new", s.handleCreatePostPOST)
	e.GET("/delete/:id", s.handleDeletePostGET)

	port := "8080"
	log.Fatal(e.Start(":" + port))
}
