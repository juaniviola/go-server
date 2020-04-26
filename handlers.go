package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// server methods
func getData(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks.items)
}

func addData(c echo.Context) error {
	t := new(Task)
	if err := c.Bind(t); err != nil {
		return err
	}

	tasks.addToSlice(*t)
	return c.JSON(http.StatusCreated, t)
}

func updateData(c echo.Context) error {
	t := echo.Map{}
	if err := c.Bind(&t); err != nil {
		return err
	}

	task := &Task{
		ID:          t["id"].(string),
		Name:        t["name"].(string),
		TaskName:    t["taskname"].(string),
		Description: t["description"].(string),
	}

	tasks.updateToSlice(task.ID, *task)
	return c.JSON(http.StatusOK, task)
}

func deleteData(c echo.Context) error {
	t := echo.Map{}
	if err := c.Bind(&t); err != nil {
		return err
	}

	tasks.removeToSlice(t["id"].(string))
	return c.String(http.StatusOK, "Deleted")
}
