package di

import "github.com/hiroyky/famiphoto/services"

func NewAuthService() services.OAuthService {
	return services.NewOAuthService(newPasswordService(), NewAuthAdapter())
}

func NewImageProcessService() services.ImageProcessService {
	return services.NewImageProcessService(NewPhotoStorageAdapter())
}

func NewPhotoService() services.PhotoService {
	return services.NewPhotoService(NewPhotoAdapter(), NewPhotoStorageAdapter())
}

func NewUserService() services.UserService {
	return services.NewUserService(NewUserAdapter(), newPasswordService())
}

func newPasswordService() services.PasswordService {
	return services.NewPasswordService()
}
