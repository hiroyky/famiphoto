package main

import (
	"context"
	"github.com/hiroyky/famiphoto/di"
)

func main() {
	uc := di.NewSearchUseCase()
	if err := uc.AppendAllPhotoDocuments(context.Background()); err != nil {
		panic(err)
	}
}
