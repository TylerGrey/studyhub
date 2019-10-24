package resolvers

import (
	"github.com/TylerGrey/hub-api/api/app/repo/mysql"

	"github.com/TylerGrey/hub-api/api/app/resolvers/args"
	"github.com/TylerGrey/hub-api/api/app/resolvers/model"
)

// CreateUser 유저 생성
func (r *Resolver) CreateUser(input args.CreateUserInput) (*model.User, error) {
	result := <-r.UserRepo.Create(mysql.User{
		Email:    input.Input.Email,
		Password: input.Input.Password,
		Name:     input.Input.Name,
	})
	if result.Err != nil {
		return nil, result.Err
	}

	return &model.User{
		Data: result.Data.(mysql.User),
	}, nil
}
