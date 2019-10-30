package handlers

import (
	"github.com/labstack/echo"
	"time"

	"github.com/pashukhin/pik-arenda-test-task/entity"
	"github.com/pashukhin/pik-arenda-test-task/service"

	"net/http"
	"strconv"
)

type Schedule struct {
	service *service.Schedule
}

func NewSchedule(service *service.Schedule) CrudHandler {
	return &Schedule{service}
}

func (h *Schedule) List(c echo.Context) error {
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

func (h *Schedule) Create(c echo.Context) error {
	o := &entity.Schedule{}
	if err := c.Bind(o); err != nil {
		return err
	}
	if err := h.service.Create(o); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, o)
}

func (h *Schedule) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	if o, err := h.service.Get(id); err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, o)
	}
}

func (h *Schedule) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	if o, err := h.service.Get(id); err != nil {
		return err
	} else {
		if err := c.Bind(o); err != nil {
			return err
		}
		if err := h.service.Update(o); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, o)
	}
}

func (h *Schedule) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	if o, err := h.service.Delete(id); err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, o)
	}
}
