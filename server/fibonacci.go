package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type fibonacciResponse struct {
	Sequence []int `json:"sequence"`
}

func newFibonacciResponse(seq []int) fibonacciResponse {
	return fibonacciResponse{
		Sequence: seq,
	}
}
func (r *Rest) fibonacci(c echo.Context) error {
	//get x parameter
	x, err := strconv.Atoi(c.QueryParam("x"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "x must be a number")
	}
	//get y parameter
	y, err := strconv.Atoi(c.QueryParam("y"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "y must be a number")
	}
	//get and response fibonacci sequence
	seq, err := r.Handler.GetFibonacciSequence(x, y)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, newFibonacciResponse(seq))
}
