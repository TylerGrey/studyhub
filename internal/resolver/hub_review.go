package resolver

import (
	"context"
	"strconv"

	"github.com/TylerGrey/studyhub/internal/mysql/repo"
	"github.com/graph-gophers/graphql-go"
)

// HubReview ...
type HubReview struct {
	Data repo.HubReview
}

// ID ...
func (h HubReview) ID() graphql.ID {
	id := strconv.FormatUint(h.Data.ID, 10)
	return graphql.ID(id)
}

// User ...
func (h HubReview) User(ctx context.Context) (*User, error) {
	return NewUser(ctx, NewUserArgs{
		ID: strconv.Itoa(int(h.Data.UserID)),
	})
}

// Rating ...
func (h HubReview) Rating() int32 {
	return 0
}

// Review ...
func (h HubReview) Review() string {
	return ""
}

// Images ...
func (h HubReview) Images() *[]Image {
	return &[]Image{}
}

// CreatedAt ...
func (h HubReview) CreatedAt() string {
	return h.Data.CreatedAt.String()
}

// UpdatedAt ...
func (h HubReview) UpdatedAt() string {
	return h.Data.UpdatedAt.String()
}

// DeletedAt ...
func (h HubReview) DeletedAt() *string {
	if h.Data.DeletedAt != nil {
		at := h.Data.DeletedAt.String()
		return &at
	}

	return nil
}
