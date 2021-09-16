package handler

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"strconv"
)

//GetFibonacciSequence returns all numbers, Fibonacci sequences with ordinal numbers from x to y.
func (h *Handler) GetFibonacciSequence(x int, y int) ([]int64, error) {
	var err error
	//check if the parameters are valid
	if err := h.ValidateParameters(x, y); err != nil {
		return nil, err
	}

	sequenceLength := (y - x) + 1
	sequence := make([]int64, sequenceLength)
	for i := x; i <= y; i++ {
		sequence[i-x], err = h.GetFibonacciNumber(i)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to get fibonacci number")
		}
	}
	return sequence, nil
}

func (h *Handler) GetFibonacciNumber(n int) (int64, error) {
	//try to find fibonacci number in memcache
	v, _ := h.Cache.Get(strconv.Itoa(n))
	if v != nil {
		return strconv.ParseInt(string(v.Value), 10, 64)
	}

	//if we couldn't find it, then count
	var result int64
	if n == 0 {
		result = 0
	} else if n < 2 {
		result = 1
	} else {
		var a int64 = 1
		var b int64 = 1
		for i := 2; i < n; i++ {
			b = a + b
			a = b - a
		}
		result = b
	}

	//write a number to cache
	err := h.Cache.Set(&memcache.Item{Key: strconv.Itoa(n), Value: []byte(strconv.FormatInt(result, 10))})
	if err != nil {
		log.Warn(errors.Wrap(err, "Failed to send the result to memcache"))
	}

	return result, nil
}

//ValidateParameters checks if the parameters are in the range of valid values
func (h *Handler) ValidateParameters(x int, y int) error {
	if x < 0 || y < 0 {
		return errors.New("Parameters cannot be less than 0")
	}
	if x > y {
		return errors.New("y must be greater than or equal to x")
	}
	if x > 92 || y > 92 {
		return errors.New("Parameters should not be more than 92")
	}
	return nil
}
