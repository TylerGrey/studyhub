package loader

import (
	"context"
	"errors"

	"github.com/TylerGrey/studyhub/internal/mysql/repo"
	"github.com/nicksrandall/dataloader"
)

// LoadUser ...
func LoadUser(ctx context.Context, id string) (repo.User, error) {
	var user repo.User

	ldr, err := extract(ctx, userLoaderKey)
	if err != nil {
		return user, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(id))()
	if err != nil {
		return user, err
	}

	user, ok := data.(repo.User)
	if !ok {
		return user, errors.New("type error")
	}

	return user, nil
}

type userGetter interface {
	Users(ctx context.Context, ids []string) (map[string]repo.User, error)
}

type userLoader struct {
	get userGetter
}

func newUserLoader(client userGetter) dataloader.BatchFunc {
	return userLoader{get: client}.loadBatch
}

func (ldr userLoader) loadBatch(ctx context.Context, ids dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(ids)
		results = make([]*dataloader.Result, n)
	)

	users, err := ldr.get.Users(ctx, ids.Keys())
	if err != nil {
		for idx := 0; idx < len(ids); idx++ {
			results[idx] = &dataloader.Result{Data: nil, Error: err}
		}

		return results
	}

	for idx, id := range ids.Keys() {
		u, ok := users[id]
		if !ok {
			results[idx] = &dataloader.Result{Data: repo.User{}, Error: errors.New("User not found")}
		}
		results[idx] = &dataloader.Result{Data: u, Error: nil}
	}

	return results
}
