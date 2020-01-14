package resolver

import (
	"context"
	"errors"
	"strconv"

	"github.com/TylerGrey/studyhub/api/app/loader"
	"github.com/TylerGrey/studyhub/internal/mysql/repo"
	"github.com/graph-gophers/graphql-go"
)

// User ...
type User struct {
	Data repo.User
}

// NewUserArgs ...
type NewUserArgs struct {
	User repo.User
	ID   string
}

// NewUser ...
func NewUser(ctx context.Context, args NewUserArgs) (*User, error) {
	var (
		user repo.User
		err  error
	)

	switch {
	case args.User.ID != 0:
		user = args.User
	case args.ID != "":
		user, err = loader.LoadUser(ctx, args.ID)
	default:
		err = errors.New("Error user resolver")
	}

	if err != nil {
		return nil, err
	}

	return &User{Data: user}, nil
}

// ID ...
func (u User) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(u.Data.ID)))
}

// Email ...
func (u User) Email() string {
	return u.Data.Email
}

// FirstName ...
func (u User) FirstName() *string {
	return u.Data.FirstName
}

// LastName ...
func (u User) LastName() *string {
	return u.Data.LastName
}

// Nickname ...
func (u User) Nickname() string {
	return u.Data.Nickname
}

// Mobile ...
func (u User) Mobile() string {
	return u.Data.Mobile
}

// Birth ...
func (u User) Birth() *string {
	return u.Data.Birth
}

// Gender ...
func (u User) Gender() *string {
	return u.Data.Gender
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
