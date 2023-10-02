package di

import "github.com/hiroyky/famiphoto/usecases"

func NewOAuthUseCase() usecases.OauthUseCase {
	return usecases.NewOauthUseCase(NewAuthService(), NewUserService())
}

func NewPhotoImportUseCase() usecases.PhotoImportUseCase {
	return usecases.NewPhotoImportUseCase(NewPhotoService(), NewImageProcessService(), NewPhotoAdapter(), NewPhotoStorageAdapter(), NewSearchAdapter(), NewUserAdapter())
}

func NewSearchUseCase() usecases.SearchUseCase {
	return usecases.NewSearchUseCase(NewSearchAdapter(), NewPhotoAdapter())
}

func NewUserUseCase() usecases.UserUseCase {
	return usecases.NewUserUseCase(NewUserAdapter(), NewUserService(), NewAuthService(), newPasswordService())
}

func NewPhotoUseCase() usecases.PhotoUseCase {
	return usecases.NewPhotoUseCase(NewPhotoAdapter())
}

func NewDownloadUseCase() usecases.DownloadUseCase {
	return usecases.NewDownloadUseCase(NewPhotoAdapter(), NewPhotoStorageAdapter())
}
