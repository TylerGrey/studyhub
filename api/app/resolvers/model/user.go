package model

import (
	"strconv"

	"github.com/TylerGrey/hub-api/api/app/repo/mysql"
	"github.com/graph-gophers/graphql-go"
)

// User ...
type User struct {
	Data mysql.User
}

// ID ...
func (u User) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(u.Data.ID)))
}

// Email ...
func (u User) Email() string {
	return u.Data.Email
}

// Name ...
func (u User) Name() string {
	return u.Data.Name
}

// CreatedAt ...
func (u User) CreatedAt() string {
	return u.Data.CreatedAt.String()
}

// UpdatedAt ...
func (u User) UpdatedAt() string {
	return u.Data.UpdatedAt.String()
}

// DeletedAt ...
func (u User) DeletedAt() *string {
	if u.Data.DeletedAt != nil {
		at := u.Data.DeletedAt.String()
		return &at
	}

	return nil
}
