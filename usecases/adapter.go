package usecases

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/hiroyky/famiphoto/entities"
	"time"
)

type RandomService interface {
	GenerateRandomString(length int) string
}

type UserAdapter interface {
	GetUser(ctx context.Context, userID string) (*entities.User, error)
	GetUsers(ctx context.Context, filter *UserFilter, limit, offset int) (entities.UserList, error)
	CountUsers(ctx context.Context, filter *UserFilter) (int, error)
	ExistUser(ctx context.Context, userID string) (bool, error)
	CreateUser(ctx context.Context, user *entities.User, password string, isInitializedPassword bool, now time.Time) (*entities.User, error)
}

type UserFilter struct {
	UserID *string
}

type UserService interface {
	AuthUserPassword(ctx context.Context, userID, password string) error
}

type UserPasswordAdapter interface {
	GetUserPassword(ctx context.Context, userID string) (*entities.UserPassword, error)
}

type AuthService interface {
	PublishUserAccessToken(ctx context.Context, client *entities.OauthClient, userID string) (string, int64, error)
	PublishCCAccessToken(ctx context.Context, client *entities.OauthClient) (string, int64, error)
	GetSession(ctx context.Context, accessToken string) (*entities.OauthSession, error)
	AuthByRefreshToken(ctx context.Context, clientID, refreshToken string) (*entities.UserAuth, error)
	UpsertUserAuth(ctx context.Context, clientID, userID string, now time.Time) (string, error)
	AuthCode(ctx context.Context, client *entities.OauthClient, code, redirectURL string) (*entities.OAuthCode, error)
	PublishAuthCode(ctx context.Context, clientID, userID, redirectURL string) (string, error)
	AuthClient(ctx context.Context, clientID, clientSecret string) (*entities.OauthClient, error)
	CreateClient(ctx context.Context, client *entities.OauthClient) (*entities.OauthClient, string, error)
	ValidateToCreateClient(ctx context.Context, client *entities.OauthClient) error
	GetUserClient(ctx context.Context, clientID string) (*entities.OauthClient, error)
	ValidateRedirectURL(ctx context.Context, clientID, redirectURL string) error
}

type PasswordService interface {
	HashPassword(password string) (string, error)
	MatchPassword(password string, hash string) (bool, error)
	GeneratePassword(length int) (string, error)
}

type OauthClientAdapter interface {
	GetByOauthClientID(ctx context.Context, id string) (*entities.OauthClient, error)
	CreateOAuthClient(ctx context.Context, client *entities.OauthClient, clientSecret string) (*entities.OauthClient, error)
	ExistOauthClient(ctx context.Context, id string) (bool, error)
}

type OauthClientRedirectURLAdapter interface {
	GetOAuthClientRedirectURLsByOAuthClientID(ctx context.Context, oauthClientID string) (entities.OAuthClientRedirectURLList, error)
	CreateOAuthClientRedirectURL(ctx context.Context, url *entities.OAuthClientRedirectURL) (*entities.OAuthClientRedirectURL, error)
}

type OauthAccessTokenAdapter interface {
	SetClientCredentialAccessToken(ctx context.Context, clientID, accessToken string, expireAt int64) error
	SetUserAccessToken(ctx context.Context, clientID, userID, accessToken string, scope entities.OauthScope, expireIn int64) error
	GetSession(ctx context.Context, accessToken string) (*entities.OauthSession, error)
}

type OauthCodeAdapter interface {
	SetCode(ctx context.Context, code *entities.OAuthCode) error
	GetCode(ctx context.Context, code string) (*entities.OAuthCode, error)
}

type UserAuthAdapter interface {
	UpsertUserAuth(ctx context.Context, m *entities.UserAuth) (*entities.UserAuth, error)
	GetUserAuth(ctx context.Context, userID, clientID string) (*entities.UserAuth, error)
	GetUserAuthByRefreshToken(ctx context.Context, refreshToken string) (*entities.UserAuth, error)
	DeleteUserAuth(ctx context.Context, userID, clientID string) error
	DeleteClientAllAuth(ctx context.Context, clientID string) error
}

type PhotoStorageAdapter interface {
	FindDirContents(dirPath string) ([]*entities.StorageFileInfo, error)
	LoadContent(path string) (entities.StorageFileData, error)
	ParsePhotoMeta(path string) (entities.PhotoMeta, error)
}

type PhotoAdapter interface {
	InsertPhoto(ctx context.Context, photo *entities.Photo) (*entities.Photo, error)
	UpdatePhoto(ctx context.Context, photo *entities.Photo) (*entities.Photo, error)
	GetPhotos(ctx context.Context, limit, offset int64) (entities.PhotoList, error)
	CountPhotos(ctx context.Context) (int64, error)
	GetPhotoByFilePath(ctx context.Context, filePath string) (*entities.Photo, error)
	GetPhotoFileByFilePath(ctx context.Context, filePath string) (*entities.PhotoFile, error)
	InsertPhotoFile(ctx context.Context, file *entities.PhotoFile) (*entities.PhotoFile, error)
	UpdatePhotoFile(ctx context.Context, file *entities.PhotoFile) (*entities.PhotoFile, error)
	GetPhotoFilesByPhotoIDs(ctx context.Context, photoIDs []int64) ([]*entities.PhotoFile, error)
	InsertPhotoMetaItem(ctx context.Context, photoID int64, meta *entities.PhotoMetaItem) (*entities.PhotoMetaItem, error)
	UpdatePhotoMetaItem(ctx context.Context, photoID int64, meta *entities.PhotoMetaItem) (*entities.PhotoMetaItem, error)
	GetPhotoMetaItemByTagID(ctx context.Context, photoID, tagID int64) (*entities.PhotoMetaItem, error)
}

type PhotoService interface {
	RegisterPhoto(ctx context.Context, filePath, fileHash, ownerID, groupID string) (*entities.PhotoFile, error)
}

type ImageProcessService interface {
	CreateThumbnails(ctx context.Context, photoFile *entities.PhotoFile, data []byte) error
}

type PhotoThumbnailAdapter interface {
	SavePreviewThumbnail(ctx context.Context, photoID int64, data []byte, groupID, ownerID string) error
}

type SearchAdapter interface {
	BulkInsertPhoto(ctx context.Context, photos []*entities.Photo, photoFiles entities.PhotoFileList, dateTimeOriginal *entities.PhotoMetaItem) (*esutil.BulkIndexerStats, error)
}
