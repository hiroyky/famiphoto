package di

import "github.com/hiroyky/famiphoto/usecases"

func NewOAuthUseCase() usecases.OauthUseCase {
	return usecases.NewOauthUseCase(NewAuthService(), NewUserService())
}

func NewPhotoImportUseCase() usecases.PhotoImportUseCase {
	return usecases.NewPhotoImportUseCase(NewPhotoService(), NewImageProcessService(), NewPhotoAdapter(), NewPhotoStorageAdapter(), NewSearchAdapter(), NewUserAdapter(), NewGroupAdapter())
}

func NewSearchUseCase() usecases.SearchUseCase {
	return usecases.NewSearchUseCase(NewSearchAdapter(), NewPhotoAdapter(), NewGroupAdapter())
}

func NewUserUseCase() usecases.UserUseCase {
	return usecases.NewUserUseCase(NewUserAdapter(), NewGroupAdapter(), NewUserService(), NewAuthService(), newPasswordService())
}

func NewGroupUseCase() usecases.GroupUseCase {
	return usecases.NewGroupUseCase(NewGroupAdapter())
}

func NewPhotoUseCase() usecases.PhotoUseCase {
	return usecases.NewPhotoUseCase(NewPhotoAdapter())
}

func NewDownloadUseCase() usecases.DownloadUseCase {
	return usecases.NewDownloadUseCase(NewPhotoAdapter(), NewGroupAdapter(), NewPhotoStorageAdapter())
}
