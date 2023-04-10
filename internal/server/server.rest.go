package server

import "github.com/labstack/echo/v4"

type Handler interface {
	Route(rest *echo.Echo)
}

func StartRest(handlers ...Handler) error {
	rest := echo.New()

	for _, handler := range handlers {
		handler.Route(rest)
	}

	return rest.Start(":7070")
}
