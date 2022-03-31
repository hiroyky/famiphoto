package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAuthHeader_Bearer(t *testing.T) {
	header := "Bearer tokentoken"
	actual1, actual2 := ParseAuthHeader(header, "Bearer")
	assert.Equal(t, "tokentoken", actual1)
	assert.Equal(t, true, actual2)
}

func TestParseAuthHeader_bearer(t *testing.T) {
	header := "Bearer tokentoken"
	actual1, actual2 := ParseAuthHeader(header, "bearer")
	assert.Equal(t, "tokentoken", actual1)
	assert.Equal(t, true, actual2)
}

func TestParseAuthHeader_empty(t *testing.T) {
	header := ""
	actual1, actual2 := ParseAuthHeader(header, "bearer")
	assert.Equal(t, "", actual1)
	assert.Equal(t, false, actual2)
}

func TestParseAuthHeader_異なるタイプ(t *testing.T) {
	header := "Basic tokentoken"
	actual1, actual2 := ParseAuthHeader(header, "Bearer")
	assert.Equal(t, "", actual1)
	assert.Equal(t, false, actual2)
}
