package server

import (
	"github.com/Nikita99100/Fibonacci/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Rest struct {
	Router  *echo.Echo
	Handler *handler.Handler
}

func (r *Rest) Route() {
	r.Router.GET("/", r.hello)
	r.Router.GET("/fibonacci", r.fibonacci)
}

func (r *Rest) hello(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>Hello</h1>")
}
