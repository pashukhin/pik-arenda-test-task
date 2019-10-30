package handlers

import (
	"github.com/labstack/echo"
)

const layout = `2006-01-02T15:04:05.000Z`

type CrudHandler interface {
	List(c echo.Context) error
	Create(c echo.Context) error
	Get(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
