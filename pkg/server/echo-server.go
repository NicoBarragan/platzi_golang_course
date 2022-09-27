package language

import (
	"net/http"

	"github.com/labstack/echo"
)

func echoServer() {
	// Instance
	e := echo.New()

	// Route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
