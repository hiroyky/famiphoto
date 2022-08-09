package di

import "github.com/hiroyky/famiphoto/usecases"

func NewOAuthUseCase() usecases.OauthUseCase {
	return usecases.NewOauthUseCase(NewAuthService(), NewUserService())
}

func NewPhotoImportUseCase() usecases.PhotoImportUseCase {
	return usecases.NewPhotoImportUseCase(NewPhotoService(), NewImageProcessService(), NewPhotoStorageAdapter())
}

func NewSearchUseCase() usecases.SearchUseCase {
	return usecases.NewSearchUseCase(NewSearchAdapter(), NewPhotoAdapter())
}

func NewUserUseCase() usecases.UserUseCase {
	return usecases.NewUserUseCase(NewUserAdapter(), newPasswordService())
}

func NewGroupUseCase() usecases.GroupUseCase {
	return usecases.NewGroupUseCase(NewGroupAdapter())
}

func NewPhotoUseCase() usecases.PhotoUseCase {
	return usecases.NewPhotoUseCase(NewPhotoAdapter())
}
