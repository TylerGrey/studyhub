package client

import (
	"context"
	"strconv"

	"github.com/TylerGrey/studyhub/internal/mysql/repo"
)

// Users ...
func (c Client) Users(ctx context.Context, ids []string) (map[string]repo.User, error) {
	userIds := []uint64{}

	for _, id := range ids {
		userID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return nil, err
		}

		userIds = append(userIds, userID)
	}

	usersMap := map[string]repo.User{}
	users, err := c.UserRepo.FindByIDs(userIds)
	for _, u := range users {
		id := strconv.Itoa(int(u.ID))
		usersMap[id] = *u
	}

	return usersMap, err
}
