package main

import (
	"net/http"

	"github.com/fahrettinEmre/todo-case/controllers"
	"github.com/fahrettinEmre/todo-case/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	models.Init()
	e := echo.New()
	e.GET("/hello", hello)
	e.GET("/getall", controllers.GetAll)
	e.POST("/create", controllers.Create)
	e.PUT("/update", controllers.Update)
	e.POST("/delete/:id", controllers.Delete)

	e.Logger.Fatal(e.Start(":8080"))

}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
