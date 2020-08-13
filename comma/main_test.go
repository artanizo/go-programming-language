package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComma(t *testing.T) {
	assert.Equal(t, "", comma(""))
	assert.Equal(t, "123", comma("123"))
	assert.Equal(t, "123,4", comma("1234"))
	assert.Equal(t, "123,45", comma("12345"))
	assert.Equal(t, "123,456", comma("123456"))
	assert.Equal(t, "123,456,789", comma("123456789"))
	assert.Equal(t, "123,456,789.44", comma("123456789.44"))
	assert.Equal(t, "-123,456,789.44", comma("-123456789.44"))
}
