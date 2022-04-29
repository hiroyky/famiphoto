package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPhoto_FileNameHash(t *testing.T) {
	p := &Photo{
		FilePath: "/photos/group-1/user-1/2022-01-4/photo-001.jpg",
	}
	actual := p.FileNameHash()
	expected := "2ce153960f310afb6484d49072d58734"
	assert.Equal(t, expected, actual)
}
