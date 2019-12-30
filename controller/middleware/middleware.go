package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/freitzzz/iped/controller/db"
	"github.com/labstack/echo"
)

// NotFoundHandler sets the NotFoundHandler of Echo's context to return 404 upon non-existent routes
func NotFoundHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusNotFound)
	}
}

// DbAccessMiddleware is a middleware that sets the database connection object using "db" as key
func DbAccessMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db.Db)
			next(c)
			return nil
		}
	}
}

// ResourceIdentifierValidationMiddleware is a middleware that checks if resource identifiers passed as params (eg: :id)
// is a valid integer or if its greater than zero
// If not it automatically responds with 404 Not Found
func ResourceIdentifierValidationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			params := c.ParamNames()
			for _, param := range params {
				if strings.Contains(param, "id") {
					id, err := strconv.Atoi(c.Param(param))
					if err != nil || id <= 0 {
						return c.NoContent(http.StatusNotFound)
					}
				}
			}
			next(c)
			return nil
		}
	}
}
