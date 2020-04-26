package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

// global variable: slice of struct Task
var tasks Tasks

// main
func main() {
  e := echo.New()
  e.Use(middleware.Logger())
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{echo.GET, echo.POST},
  }))

  e.GET("/", getData)
  e.POST("/add", addData)
  e.POST("/update", updateData)
  e.POST("/delete", deleteData)
  e.Logger.Fatal(e.Start(":1323"))
}
