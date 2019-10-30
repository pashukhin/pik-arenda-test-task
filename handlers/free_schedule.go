package handlers

import (
	"net/http"
	"github.com/labstack/echo"
	"time"

	"github.com/pashukhin/pik-arenda-test-task/service"
)

type FreeSchedule struct {
	service *service.FreeSchedule
}

func NewFreeSchedule(service *service.FreeSchedule) CrudHandler {
	return &FreeSchedule{service}
}

func (h *FreeSchedule) List(c echo.Context) error {
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
	if list, err := h.service.List(from, to); err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, list)
	}
}

func (h *FreeSchedule) Create(c echo.Context) error {
	return c.NoContent(http.StatusMethodNotAllowed)
}

func (h *FreeSchedule) Get(c echo.Context) error {
	return c.NoContent(http.StatusMethodNotAllowed)
}

func (h *FreeSchedule) Update(c echo.Context) error {
	return c.NoContent(http.StatusMethodNotAllowed)
}

func (h *FreeSchedule) Delete(c echo.Context) error {
	return c.NoContent(http.StatusMethodNotAllowed)
}
