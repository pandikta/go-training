package main_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Pembagian(a int, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("ga bs dibago 0 ")
	}

	return a / b, nil
}

func TestFunction(t *testing.T) {
	bagi, err := Pembagian(8, 0)
	assert.Equal(t, -1, bagi, err)

	// bagi2 := Pembagian(9, 2)
	// assert.Equal(t, 4.5, bagi2)
}
