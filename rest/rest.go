package rest

import (
	"github.com/Nikita99100/Fibonacci/handler"
	"github.com/labstack/echo/v4"
)

type Rest struct {
	Router  *echo.Echo
	Handler handler.Handler
}
