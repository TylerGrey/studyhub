package resolvers

import (
	"github.com/TylerGrey/study-hub/internal/mysql/repo"

	"github.com/TylerGrey/study-hub/api/app/resolvers/args"
	"github.com/TylerGrey/study-hub/internal/resolver"
)

// CreateUser 유저 생성
func (r *Resolver) CreateUser(input args.CreateUserInput) (*resolver.User, error) {
	user, err := r.UserRepo.Create(repo.User{
		Email:    input.Input.Email,
		Password: input.Input.Password,
		Name:     input.Input.Name,
	})
	if err != nil {
		return nil, err
	}

	return &resolver.User{
		Data: *user,
	}, nil
}
