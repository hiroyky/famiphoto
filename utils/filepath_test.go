package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirectories(t *testing.T) {
	actual := Directories("/usr/local/bin/emacs")
	expected := []string{"usr", "local", "bin"}
	assert.Equal(t, expected, actual)
}

func TestDirectories_カレントディレクトリ(t *testing.T) {
	actual := Directories("./emacs")
	expected := []string{}
	assert.Equal(t, expected, actual)
}
