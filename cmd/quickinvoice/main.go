package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zivlakmilos/quickinvoice/pkg/quickinvoice"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.POST("/invoice", func(c echo.Context) error {
		data, err := quickinvoice.DecodeJson(c.Request().Body)
		if err != nil {
			return err
		}
		err = quickinvoice.GenerateInvoice(data, c.Response())
		return err
	})

	e.Logger.Fatal(e.Start(":8080"))
}
