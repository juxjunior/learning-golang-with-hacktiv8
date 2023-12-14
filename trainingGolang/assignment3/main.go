package main

import (
	"assignment3/service"
	"assignment3/util"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/update", service.UpdateData)

	go util.MakeRequests()

	e.Logger.Fatal(e.Start(":8888"))
}
