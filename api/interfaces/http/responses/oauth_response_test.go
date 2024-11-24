package responses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOAuthCodeRedirectURL(t *testing.T) {
	actual, err := NewOAuthCodeRedirectURL("https://famiphoto.jp/login", "code1", "state1")
	expected := "https://famiphoto.jp/login?code=code1&state=state1"
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
