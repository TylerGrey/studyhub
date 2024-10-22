package resolvers

import (
	"strconv"

	"github.com/TylerGrey/studyhub/internal/resolver"
)

// User 유저 조회
func (r *Resolver) User(args struct {
	ID string
}) (*resolver.User, error) {
	id, err := strconv.ParseUint(args.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	user, err := r.UserRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &resolver.User{
		Data: *user,
	}, nil
}
