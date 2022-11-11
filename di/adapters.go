package di

import (
	"github.com/hiroyky/famiphoto/drivers/es"
	"github.com/hiroyky/famiphoto/infrastructures"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
)

func NewAuthAdapter() infrastructures.AuthAdapter {
	return infrastructures.NewAuthAdapter(
		newOAuthClientRepository(),
		NewOauthAccessTokenRepository(),
		NewOauthClientRedirectURLRepository(),
		newOauthCodeRepository(),
		newUserAuthRepository(),
	)
}

func newOAuthClientRepository() repositories.OAuthClientRepository {
	return repositories.NewOAuthClientRepository(NewMySQLDriver())
}

func NewOauthAccessTokenRepository() repositories.OauthAccessTokenRepository {
	return repositories.NewOauthAccessTokenRepository(NewOAuthRedisDB())
}

func NewOauthClientRedirectURLRepository() repositories.OAuthClientRedirectURLRepository {
	return repositories.NewOauthClientRedirectURLRepository(NewMySQLDriver())
}

func newOauthCodeRepository() repositories.OauthCodeRepository {
	return repositories.NewOauthCodeRepository(NewOAuthRedisDB())
}

func newUserAuthRepository() repositories.UserAuthRepository {
	return repositories.NewUserAuthRepository(NewMySQLDriver())
}

func NewPhotoAdapter() infrastructures.PhotoAdapter {
	return infrastructures.NewPhotoAdapter(newPhotoRepository(), newPhotoFileRepository(), newExifRepository())
}

func newPhotoRepository() repositories.PhotoRepository {
	return repositories.NewPhotoRepository(NewMySQLDriver())
}

func newPhotoFileRepository() repositories.PhotoFileRepository {
	return repositories.NewPhotoFileRepository(NewMySQLDriver())
}

func newExifRepository() repositories.ExifRepository {
	return repositories.NewExifRepository(NewMySQLDriver())
}

func NewPhotoStorageAdapter() infrastructures.PhotoStorageAdapter {
	return infrastructures.NewPhotoStorageAdapter(newPhotoStorageRepository(), newPhotoThumbnailRepository())
}

func newPhotoThumbnailRepository() repositories.PhotoThumbnailRepository {
	return repositories.NewPhotoThumbnailRepository(NewPhotoThumbnailStorageDriver(), NewMySQLDriver())
}

func newPhotoStorageRepository() repositories.PhotoStorageRepository {
	return repositories.NewPhotoStorageRepository(NewMediaSambaStorageDriver())
}

func NewSearchAdapter() infrastructures.SearchAdapter {
	return infrastructures.NewSearchAdapter(newElasticSearchRepo())
}

func newElasticSearchRepo() repositories.ElasticSearchRepository {
	return repositories.NewElasticSearchRepository(es.NewSearchClient(), es.NewBulkClient)
}

func NewUserAdapter() infrastructures.UserAdapter {
	return infrastructures.NewUserAdapter(newUserRepository(), newGroupRepository(), newUserPasswordRepository())
}

func newUserRepository() repositories.UserRepository {
	return repositories.NewUserRepository(NewMySQLDriver())
}

func newUserPasswordRepository() repositories.UserPasswordRepository {
	return repositories.NewUserPasswordRepository(NewMySQLDriver())
}

func newGroupRepository() repositories.GroupRepository {
	return repositories.NewGroupRepository(NewMySQLDriver())
}

func NewGroupAdapter() infrastructures.GroupAdapter {
	return infrastructures.NewGroupAdapter(newGroupRepository())
}
