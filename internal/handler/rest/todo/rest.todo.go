package todo

import (
	"context"
	"crypto/subtle"
	"github.com/erry-az/test-go/internal/entity"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type todoRepo interface {
	Get(ctx context.Context, id uint) (entity.Todo, error)
	Create(ctx context.Context, todo entity.Todo) (uint, error)
}

type Handler struct {
	repo todoRepo
}

func New(repo todoRepo) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Route(rest *echo.Echo) {
	rest.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("xxxx")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("xxxx")) == 1 {
			return true, nil
		}

		return false, nil
	}))

	rest.GET("/todo/:id", h.Get)
	rest.POST("/todo", h.Create)
}

func (h *Handler) Get(c echo.Context) error {
	var request GetRequest
	err := c.Bind(&request)
	if err != nil {
		return err
	}

	todo, err := h.repo.Get(c.Request().Context(), request.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, Response[entity.Todo]{
		Data: todo,
		OK:   true,
	})
}

func (h *Handler) Create(c echo.Context) error {
	var request CreateRequest

	err := c.Bind(&request)
	if err != nil {
		return err
	}

	todoID, err := h.repo.Create(c.Request().Context(), entity.Todo{
		TaskName:    request.TaskName,
		Description: request.Description,
		IsDone:      request.IsDone,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, Response[uint]{
		Data: todoID,
		OK:   true,
	})
}
