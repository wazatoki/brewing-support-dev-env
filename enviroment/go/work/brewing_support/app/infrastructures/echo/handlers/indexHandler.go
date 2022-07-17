package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

/*
Root root handler
*/
func Root(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, c.Request().URL.Host+"/index.html")
}

/*
Index index handler
*/
func Index(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, c.Request().URL.Host+"/index.html")
}
