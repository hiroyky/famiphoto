// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Connection interface {
	IsConnection()
}

type Edge interface {
	IsEdge()
}

type Node interface {
	IsNode()
}

type Pagination interface {
	IsPagination()
}

type CreateGroupInput struct {
	GroupID string `json:"groupId"`
	Name    string `json:"name"`
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

type GroupEdge struct {
	Cursor string `json:"cursor"`
	Node   *Group `json:"node"`
}

func (GroupEdge) IsEdge() {}

type GroupPagination struct {
	PageInfo *PaginationInfo `json:"pageInfo"`
	Nodes    []*Group        `json:"nodes"`
}

func (GroupPagination) IsPagination() {}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

type PaginationInfo struct {
	Page             int  `json:"page"`
	PaginationLength int  `json:"paginationLength"`
	HasNextPage      bool `json:"hasNextPage"`
	HasPreviousPage  bool `json:"hasPreviousPage"`
	Count            int  `json:"count"`
	TotalCount       int  `json:"totalCount"`
}

type UserEdge struct {
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
}

func (UserEdge) IsEdge() {}

type UserPagination struct {
	PageInfo *PaginationInfo `json:"pageInfo"`
	Nodes    []*User         `json:"nodes"`
}

func (UserPagination) IsPagination() {}

type UserPassword struct {
	ID            string `json:"id"`
	LastModified  string `json:"lastModified"`
	IsInitialized bool   `json:"isInitialized"`
}

func (UserPassword) IsNode() {}

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

func (e *OauthClientScope) UnmarshalGQL(v interface{}) error {
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

func (e *OauthClientType) UnmarshalGQL(v interface{}) error {
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

func (e *UserStatus) UnmarshalGQL(v interface{}) error {
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
