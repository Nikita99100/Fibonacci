package handler

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/pkg/errors"
)

type Handler struct {
	Cache *memcache.Client
}

func (h *Handler) GetFibonacciSequence(x int, y int) ([]int, error) {
	var err error
	sequenceLength := (y - x) + 1
	sequence := make([]int, sequenceLength)
	for i := x; i <= y; i++ {
		sequence[i-x], err = h.GetFibonacciNumber(i)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to get fibonacci number")
		}
	}
	return sequence, nil
}

func (h *Handler) GetFibonacciNumber(n int) (int, error) {
	if n == 0 {
		return 0, nil
	} else if n < 2 {
		return 1, nil
	}
	var a int = 1
	var b int = 1
	for i := 2; i < n; i++ {
		b = a + b
		a = b - a
	}
	return b, nil
}
