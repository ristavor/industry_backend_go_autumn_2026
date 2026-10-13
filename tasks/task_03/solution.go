package main

import (
	"errors"
	"strconv"
)

var ErrZero = errors.New("zero is not allowed")

func fizzBuzz(n int) (string, error) {
	if n == 0 {
		return "", ErrZero
	}
	var s string
	if n%3 == 0 {
		s += "Fizz"
	}
	if n%5 == 0 {
		s += "Buzz"
	}
	if s == "" {
		s = strconv.Itoa(n)
	}
	return s, nil
}
