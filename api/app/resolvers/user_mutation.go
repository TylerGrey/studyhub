package resolvers

import (
	"github.com/TylerGrey/studyhub/internal/mysql/repo"
	"github.com/google/uuid"

	"github.com/TylerGrey/studyhub/api/app/resolvers/args"
	"github.com/TylerGrey/studyhub/internal/resolver"
)

// CreateUser 유저 생성
func (r *Resolver) CreateUser(input args.CreateUserInput) (*resolver.User, error) {
	user, err := r.UserRepo.Create(repo.User{
		UUID:      uuid.New().String(),
		Email:     input.Input.Email,
		Password:  input.Input.Password,
		FirstName: input.Input.FirstName,
		LastName:  input.Input.LastName,
		Nickname:  input.Input.Nickname,
		Mobile:    input.Input.Mobile,
		Birth:     input.Input.Birth,
		Gender:    input.Input.Gender,
	})
	if err != nil {
		return nil, err
	}

	return &resolver.User{
		Data: *user,
	}, nil
}
