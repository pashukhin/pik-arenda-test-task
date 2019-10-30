package handlers

import (
	"github.com/labstack/echo"
	"time"

	"github.com/pashukhin/pik-arenda-test-task/entity"
	"github.com/pashukhin/pik-arenda-test-task/service"

	"net/http"
	"strconv"
)

type Task struct {
	service *service.Task
}

func NewTask(service *service.Task) CrudHandler {
	return &Task{service}
}

func (h *Task) List(c echo.Context) error {
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

func (h *Task) Create(c echo.Context) error {
	o := &entity.Task{}
	if err := c.Bind(o); err != nil {
		return err
	}
	if err := h.service.Create(o); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, o)
}

func (h *Task) Get(c echo.Context) error {
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

func (h *Task) Update(c echo.Context) error {
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

func (h *Task) Delete(c echo.Context) error {
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
