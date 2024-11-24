package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveTrailingSlash_スラッシュがある場合(t *testing.T) {
	actual := RemoveTrailingSlash("http://localhost/")
	expected := "http://localhost"
	assert.Equal(t, expected, actual)
}

func TestRemoveTrailingSlash_スラッシュがない場合(t *testing.T) {
	actual := RemoveTrailingSlash("http://localhost")
	expected := "http://localhost"
	assert.Equal(t, expected, actual)
}
