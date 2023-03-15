package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	e.Start(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")))
}
