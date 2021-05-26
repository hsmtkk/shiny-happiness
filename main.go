package main

import (
	"github.com/hsmtkk/shiny-happiness/handle"
	"github.com/hsmtkk/shiny-happiness/repo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"log"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}
	e.Renderer = t

	repo, err := repo.NewSqliteRepo("go-data.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	handler := handle.New(repo)

	e.GET("/todo", handler.GetAll)
	e.GET("/todo/:id", handler.Get)
	e.GET("/todo/edit/:id", handler.EditGet)
	e.POST("/todo/edit/:id", handler.EditPost)
	e.GET("/todo/delete/:id", handler.DeleteGet)
	e.POST("/todo/delete/:id", handler.DeletePost)
	e.GET("/todo/new", handler.NewGet)
	e.POST("/todo/new", handler.NewPost)

	e.Logger.Fatal(e.Start(":8000"))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
