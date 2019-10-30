package handlers

import (
	"github.com/labstack/echo"
	"github.com/pashukhin/pik-arenda-test-task/service"
	"net/http"
	"time"
)

type FreeSlot struct {
	fs *service.FreeSlot
}

func NewFreeSlot(fs *service.FreeSlot) CrudHandler {
	return &FreeSlot{fs}
}

func (h *FreeSlot) List(c echo.Context) error {
	fromStr, toStr := c.QueryParam(`from`), c.QueryParam(`to`)
	var from, to *time.Time
	if fromStr != `` {
		if t, err := time.Parse(layout, fromStr); err != nil {
			return err
		} else {
			from = &t
		}
	}
	if toStr != `` {
		if t, err := time.Parse(layout, toStr); err != nil {
			return err
		} else {
			to = &t
		}
	}
	slots, err := h.fs.List(from, to)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, slots)
}

func (h *FreeSlot) Create(c echo.Context) error {
	return c.NoContent(http.StatusMethodNotAllowed)
}

func (h *FreeSlot) Get(c echo.Context) error {
	return c.NoContent(http.StatusMethodNotAllowed)
}

func (h *FreeSlot) Update(c echo.Context) error {
	return c.NoContent(http.StatusMethodNotAllowed)
}

func (h *FreeSlot) Delete(c echo.Context) error {
	return c.NoContent(http.StatusMethodNotAllowed)
}
