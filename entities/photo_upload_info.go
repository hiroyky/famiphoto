package entities

type PhotoUploadSign struct {
	SignToken string
	ExpireAt  int
}

type PhotoUploadInfo struct {
	UserID  string
	GroupID string
}
