package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitPath(t *testing.T) {
	actual := SplitPath("/usr/local/bin/emacs")
	expected := []string{"usr", "local", "bin", "emacs"}
	assert.Equal(t, expected, actual)
}

func TestSplitPath_カレントディレクトリから指定(t *testing.T) {
	actual := SplitPath("./emacs")
	expected := []string{"emacs"}
	assert.Equal(t, expected, actual)
}

func TestSplitPath_カレントディレクトリ(t *testing.T) {
	actual := SplitPath("./")
	expected := []string{}
	assert.Equal(t, expected, actual)
}

func TestFileNameExceptExt(t *testing.T) {
	p := "photo/group/user/photo1.jpg"
	actual := FileNameExceptExt(p)
	expected := "photo1"
	assert.Equal(t, expected, actual)
}

func TestFileNameExceptExt_拡張子無し(t *testing.T) {
	p := "photo/group/user/photo1"
	actual := FileNameExceptExt(p)
	expected := "photo1"
	assert.Equal(t, expected, actual)
}

func TestFileNameExceptExt_拡張子(t *testing.T) {
	p := "photo/group/user/photo1.ARW"
	actual := FileNameExceptExt(p)
	expected := "photo1"
	assert.Equal(t, expected, actual)
}
