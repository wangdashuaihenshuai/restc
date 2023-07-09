package restc

import (
	_ "embed"
	"strings"

	"github.com/labstack/echo/v4"
)

//go:embed index.html
var HTML string

func Echo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 检查 Content-Type
		acceptHeader := c.Request().Header.Get("Accept")
		if strings.Contains(acceptHeader, "text/html") {
			return c.HTML(200, HTML)
		}

		return next(c)
	}
}
