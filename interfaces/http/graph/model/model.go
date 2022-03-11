package model

type Photo struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FilePath   string `json:"filePath"`
	ImportedAt string `json:"importedAt"`
	GroupID    string `json:"groupId"`
	OwnerID    string `json:"ownerId"`
}

func (Photo) IsNode() {}

type User struct {
	ID     string     `json:"id"`
	Name   string     `json:"name"`
	Status UserStatus `json:"status"`
}

func (User) IsNode() {}

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (Group) IsNode() {}
