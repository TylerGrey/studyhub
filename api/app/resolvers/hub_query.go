package resolvers

import (
	"fmt"
	"strconv"

	"github.com/TylerGrey/studyhub/api/app/resolvers/args"
	"github.com/TylerGrey/studyhub/internal/mysql/repo"
	"github.com/TylerGrey/studyhub/internal/resolver"
)

// Hub 유저 조회
func (r *Resolver) Hub(args struct {
	ID string
}) (*resolver.Hub, error) {
	id, err := strconv.ParseUint(args.ID, 10, 64)
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

// HubReviews 허브 리뷰 조회
func (r *Resolver) HubReviews(args struct {
	HubID string
}) (*[]*resolver.HubReview, error) {
	hubID, err := strconv.ParseUint(args.HubID, 10, 64)
	if err != nil {
		return nil, err
	}

	reviews, err := r.HubReviewRepo.List(hubID)
	if err != nil {
		return nil, err
	}

	resolvers := []*resolver.HubReview{}
	for _, r := range reviews {
		resolvers = append(resolvers, &resolver.HubReview{
			Data: *r,
		})
	}

	return &resolvers, nil
}
