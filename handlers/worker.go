package handlers

import (
	"github.com/labstack/echo"

	"github.com/pashukhin/pik-arenda-test-task/entity"
	"github.com/pashukhin/pik-arenda-test-task/service"

	"net/http"
	"strconv"
)

type Worker struct {
	service *service.Worker
}

func NewWorker(service *service.Worker) CrudHandler {
	return &Worker{service}
}

func (h *Worker) List(c echo.Context) error {
	if list, err := h.service.List(); err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, list)

	}
}

func (h *Worker) Create(c echo.Context) error {
	o := &entity.Worker{}
	if err := c.Bind(o); err != nil {
		return err
	}
	if err := h.service.Create(o); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, o)
}

func (h *Worker) Get(c echo.Context) error {
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

func (h *Worker) Update(c echo.Context) error {
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

func (h *Worker) Delete(c echo.Context) error {
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
