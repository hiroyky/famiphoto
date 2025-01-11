package repositories

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/drivers/ollama"
	"github.com/hiroyky/famiphoto/utils/cast"
	"github.com/ollama/ollama/api"
)

type PhotoDescribeRepository interface {
	DescribeEn(ctx context.Context, photoData []byte) (string, error)
	TranslateToJa(ctx context.Context, text string) (string, error)
}

func NewPhotoDescribeRepository(client ollama.Client) PhotoDescribeRepository {
	return &photoDescribeRepository{
		client: client,
	}
}

type photoDescribeRepository struct {
	client ollama.Client
}

func (r *photoDescribeRepository) DescribeEn(ctx context.Context, photoData []byte) (string, error) {
	var res string
	err := r.client.Generate(ctx, &api.GenerateRequest{
		Model: "llava",
		Prompt: "Please analyze the uploaded image and provide a detailed description of its contents. Include information about the objects, people, environment, colors, and any other notable details visible in the image. Be specific and thorough." +
			"And Please describe the main subject and background details of the uploaded image. Focus on identifying objects, their arrangement, and any prominent colors or themes present. In English.",
		Suffix:    "",
		System:    "",
		Template:  "",
		Context:   nil,
		Stream:    cast.Ptr(false),
		Raw:       false,
		Format:    nil,
		KeepAlive: nil,
		Images:    []api.ImageData{photoData},
		Options:   nil,
	}, func(response api.GenerateResponse) error {
		res = response.Response
		return nil
	})

	return res, err
}

func (r *photoDescribeRepository) TranslateToJa(ctx context.Context, text string) (string, error) {
	var res string
	r.client.Generate(ctx, &api.GenerateRequest{
		Model:     "llava",
		Prompt:    fmt.Sprintf("次の文章を日本語に翻訳してください。```%s```", text),
		Suffix:    "",
		System:    "",
		Template:  "",
		Context:   nil,
		Stream:    cast.Ptr(false),
		Raw:       false,
		Format:    nil,
		KeepAlive: nil,
		Images:    nil,
		Options:   nil,
	}, func(response api.GenerateResponse) error {
		res = response.Response
		return nil
	})
	return res, nil
}
