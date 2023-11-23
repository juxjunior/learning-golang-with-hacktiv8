package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	fmt.Println("Hello, World", e)
}
