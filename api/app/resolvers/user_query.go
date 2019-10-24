package resolvers

import (
	"strconv"

	"github.com/TylerGrey/hub-api/api/app/repo/mysql"
	"github.com/TylerGrey/hub-api/api/app/resolvers/model"
)

// User 유저 조회
func (r *Resolver) User(args struct {
	ID string
}) (*model.User, error) {
	id, err := strconv.ParseInt(args.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	result := <-r.UserRepo.FindByID(id)
	if result.Err != nil {
		return nil, result.Err
	}

	return &model.User{
		Data: result.Data.(mysql.User),
	}, nil
}
