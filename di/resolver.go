package di

import "github.com/hiroyky/famiphoto/interfaces/http/graph"

func NewResolver() *graph.Resolver {
	return graph.NewResolver(NewUserUseCase(), NewPhotoUseCase(), NewSearchUseCase(), NewOAuthUseCase(), NewPhotoImportUseCase())
}
