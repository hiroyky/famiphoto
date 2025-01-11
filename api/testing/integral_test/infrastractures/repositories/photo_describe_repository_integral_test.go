package repositories

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/di"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestPhotoDescribeRepository(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Run("写真を分析する", func(t *testing.T) {
		repo := di.NewPhotoDescribeRepository()
		imageFile, err := os.ReadFile(path.Join("../../../resources/test_image001.jpg"))
		assert.NoError(t, err)

		description, err := repo.DescribeEn(context.Background(), imageFile)
		assert.NoError(t, err)
		fmt.Println(description)

		ja, err := repo.TranslateToJa(context.Background(), description)
		assert.NoError(t, err)
		fmt.Println(ja)
		t.Fail()
	})
}
