package usecases

import (
	"github.com/hiroyky/famiphoto/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPhotoImportUseCase_parseBasePath_ルートパス(t *testing.T) {
	src := "photo/group_id1/user_id1"

	uc := photoImportUseCase{}
	actual1, actual2, err := uc.parseBasePath(src)
	assert.NoError(t, err)
	expected1 := "group_id1"
	expected2 := "user_id1"
	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
}

func TestPhotoImportUseCase_parseBasePath_下位層(t *testing.T) {
	src := "photo/group_id1/user_id1/subdir1"

	uc := photoImportUseCase{}
	actual1, actual2, err := uc.parseBasePath(src)
	assert.NoError(t, err)
	expected1 := "group_id1"
	expected2 := "user_id1"
	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
}

func TestPhotoImportUseCase_parseBasePath_不正(t *testing.T) {
	src := "photo"

	uc := photoImportUseCase{}
	_, _, err := uc.parseBasePath(src)
	assert.Error(t, err)
	assert.True(t, errors.IsErrCode(err, errors.InvalidFilePathFatal))
}
