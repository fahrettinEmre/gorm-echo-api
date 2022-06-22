package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fahrettinEmre/todo-case/models"
	"github.com/labstack/echo"
)

func GetAll(c echo.Context) error {
	var todos []models.Todo
	models.DB.Find(&todos)
	fmt.Println(todos)
	return c.JSON(http.StatusOK, todos)
}

func Create(c echo.Context) error {
	t := new(models.Todo)
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	models.DB.Create(t)
	return c.JSON(http.StatusCreated, nil)
}

func Update(c echo.Context) error {

	t := new(models.Todo)
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if t.ID == 0 {
		return c.JSON(http.StatusBadRequest, nil)
	}

	models.DB.Save(t)
	return c.JSON(http.StatusOK, nil)
}

func Delete(c echo.Context) error {
	id := c.Param("id")
	t := new(models.Todo)
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	t.ID = uint(parsedId)

	models.DB.Delete(&t)

	return c.JSON(http.StatusOK, nil)
}
