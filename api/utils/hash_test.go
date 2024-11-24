package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMD5(t *testing.T) {
	src := "abcdefg"
	actual := MD5(src)
	expected := "7ac66c0f148de9519b8bd264312c4d64"
	assert.Equal(t, expected, actual)
}
