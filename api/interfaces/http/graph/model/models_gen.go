// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Connection interface {
	IsConnection()
	GetPageInfo() *PageInfo
	GetEdges() []Edge
	GetNodes() []Node
}

type Edge interface {
	IsEdge()
	GetCursor() string
	GetNode() Node
}

type Node interface {
	IsNode()
	GetID() string
}

type Pagination interface {
	IsPagination()
	GetPageInfo() *PaginationInfo
	GetNodes() []Node
}

type CreateOauthClientInput struct {
	ClientID     string           `json:"clientId"`
	Name         string           `json:"name"`
	Scope        OauthClientScope `json:"scope"`
	ClientType   OauthClientType  `json:"clientType"`
	RedirectUrls []string         `json:"redirectUrls"`
}

type CreateUserInput struct {
	UserID   string `json:"userId"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type DateAggregationItem struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Date  int `json:"date"`
	Num   int `json:"num"`
}

type GqlStatus struct {
	Status string `json:"status"`
}

type IndexingPhotosInput struct {
	Fast bool `json:"fast"`
}

type Mutation struct {
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}

type PaginationInfo struct {
	Offset           int  `json:"offset"`
	Limit            int  `json:"limit"`
	Page             int  `json:"page"`
	PaginationLength int  `json:"paginationLength"`
	HasNextPage      bool `json:"hasNextPage"`
	HasPreviousPage  bool `json:"hasPreviousPage"`
	Count            int  `json:"count"`
	TotalCount       int  `json:"totalCount"`
}

type PhotoExif struct {
	ID          string `json:"id"`
	TagID       int    `json:"tagId"`
	TagType     string `json:"tagType"`
	ValueString string `json:"valueString"`
}

func (PhotoExif) IsNode()            {}
func (this PhotoExif) GetID() string { return this.ID }

type PhotoPagination struct {
	PageInfo *PaginationInfo `json:"pageInfo"`
	Nodes    []*Photo        `json:"nodes"`
}

func (PhotoPagination) IsPagination()                     {}
func (this PhotoPagination) GetPageInfo() *PaginationInfo { return this.PageInfo }
func (this PhotoPagination) GetNodes() []Node {
	if this.Nodes == nil {
		return nil
	}
	interfaceSlice := make([]Node, 0, len(this.Nodes))
	for _, concrete := range this.Nodes {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type PhotoUploadInfo struct {
	UploadURL string `json:"uploadUrl"`
	ExpireAt  int    `json:"expireAt"`
}

type Query struct {
}

type UpdateMeInput struct {
	Name string `json:"name"`
}

type UserEdge struct {
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
}

func (UserEdge) IsEdge()                {}
func (this UserEdge) GetCursor() string { return this.Cursor }
func (this UserEdge) GetNode() Node     { return *this.Node }

type UserPagination struct {
	PageInfo *PaginationInfo `json:"pageInfo"`
	Nodes    []*User         `json:"nodes"`
}

func (UserPagination) IsPagination()                     {}
func (this UserPagination) GetPageInfo() *PaginationInfo { return this.PageInfo }
func (this UserPagination) GetNodes() []Node {
	if this.Nodes == nil {
		return nil
	}
	interfaceSlice := make([]Node, 0, len(this.Nodes))
	for _, concrete := range this.Nodes {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type UserPassword struct {
	ID            string `json:"id"`
	LastModified  string `json:"lastModified"`
	IsInitialized bool   `json:"isInitialized"`
}

func (UserPassword) IsNode()            {}
func (this UserPassword) GetID() string { return this.ID }

type OauthClientScope string

const (
	OauthClientScopeGeneral OauthClientScope = "general"
	OauthClientScopeAdmin   OauthClientScope = "admin"
)

var AllOauthClientScope = []OauthClientScope{
	OauthClientScopeGeneral,
	OauthClientScopeAdmin,
}

func (e OauthClientScope) IsValid() bool {
	switch e {
	case OauthClientScopeGeneral, OauthClientScopeAdmin:
		return true
	}
	return false
}

func (e OauthClientScope) String() string {
	return string(e)
}

func (e *OauthClientScope) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OauthClientScope(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OauthClientScope", str)
	}
	return nil
}

func (e OauthClientScope) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OauthClientType string

const (
	OauthClientTypeUserClient       OauthClientType = "UserClient"
	OauthClientTypeClientCredential OauthClientType = "ClientCredential"
)

var AllOauthClientType = []OauthClientType{
	OauthClientTypeUserClient,
	OauthClientTypeClientCredential,
}

func (e OauthClientType) IsValid() bool {
	switch e {
	case OauthClientTypeUserClient, OauthClientTypeClientCredential:
		return true
	}
	return false
}

func (e OauthClientType) String() string {
	return string(e)
}

func (e *OauthClientType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OauthClientType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OauthClientType", str)
	}
	return nil
}

func (e OauthClientType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserStatus string

const (
	UserStatusActive     UserStatus = "Active"
	UserStatusWithdrawal UserStatus = "Withdrawal"
)

var AllUserStatus = []UserStatus{
	UserStatusActive,
	UserStatusWithdrawal,
}

func (e UserStatus) IsValid() bool {
	switch e {
	case UserStatusActive, UserStatusWithdrawal:
		return true
	}
	return false
}

func (e UserStatus) String() string {
	return string(e)
}

func (e *UserStatus) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserStatus", str)
	}
	return nil
}

func (e UserStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
