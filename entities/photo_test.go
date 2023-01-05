package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPhotoFile_FileType_JPG(t *testing.T) {
	photoFile := &PhotoFile{
		PhotoFileID: 0,
		PhotoID:     0,
		FilePath:    "/gid/uid/test1.JPG",
		ImportedAt:  time.Time{},
		GroupID:     "",
		OwnerID:     "",
		FileHash:    "",
	}
	actual := photoFile.FileType()
	expected := PhotoFileTypeJPEG
	assert.Equal(t, expected, actual)
}
