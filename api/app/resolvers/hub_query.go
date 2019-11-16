package resolvers

import (
	"fmt"
	"strconv"

	"github.com/TylerGrey/study-hub/api/app/resolvers/args"
	"github.com/TylerGrey/study-hub/internal/mysql/repo"
	"github.com/TylerGrey/study-hub/internal/resolver"
)

// Hub 유저 조회
func (r *Resolver) Hub(args struct {
	ID string
}) (*resolver.Hub, error) {
	id, err := strconv.ParseInt(args.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	hub, err := r.HubRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &resolver.Hub{
		Data: *hub,
	}, nil
}

// Hubs 유저 조회
func (r *Resolver) Hubs(args args.HubsArgs) (resolver.HubConnection, error) {
	if args.First == nil && args.Last == nil {
		return resolver.HubConnection{}, fmt.Errorf("You must provide a `first` or `last` value to properly paginate the `hubs` connection")
	} else if args.First != nil && args.Last != nil {
		return resolver.HubConnection{}, fmt.Errorf("Passing both `first` and `last` to paginate the `hubs` connection is not supported")
	}

	hubsArgs := repo.ListArgs{
		First:  args.First,
		Last:   args.Last,
		After:  args.After,
		Before: args.Before,
	}
	if args.OrderBy != nil {
		hubsArgs.Order = &repo.Order{
			Field:     args.OrderBy.Field,
			Direction: args.OrderBy.Direction,
		}
	}
	hubs, page, err := r.HubRepo.List(hubsArgs)
	if err != nil {
		return resolver.HubConnection{}, err
	}

	return resolver.HubConnection{
		Data: hubs,
		Page: page,
	}, nil
}
