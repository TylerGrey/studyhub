package resolver

import (
	"strconv"

	"github.com/TylerGrey/study-hub/internal/mysql/repo"
	"github.com/graph-gophers/graphql-go"
)

// HubIncorrectInfo ...
type HubIncorrectInfo struct {
	Data repo.HubIncorrectInfo
}

// ID ...
func (h HubIncorrectInfo) ID() graphql.ID {
	id := strconv.FormatUint(h.Data.ID, 10)
	return graphql.ID(id)
}

// HubID ...
func (h HubIncorrectInfo) HubID() graphql.ID {
	id := strconv.FormatUint(h.Data.HubID, 10)
	return graphql.ID(id)
}

// UserID ...
func (h HubIncorrectInfo) UserID() graphql.ID {
	id := strconv.FormatUint(h.Data.UserID, 10)
	return graphql.ID(id)
}

// Message ...
func (h HubIncorrectInfo) Message() string {
	return h.Data.Message
}

// IsFixed ...
func (h HubIncorrectInfo) IsFixed() bool {
	return h.Data.IsFixed
}

// CreatedAt ...
func (h HubIncorrectInfo) CreatedAt() string {
	return h.Data.CreatedAt.String()
}

// UpdatedAt ...
func (h HubIncorrectInfo) UpdatedAt() string {
	return h.Data.UpdatedAt.String()
}

// DeletedAt ...
func (h HubIncorrectInfo) DeletedAt() *string {
	if h.Data.DeletedAt != nil {
		at := h.Data.DeletedAt.String()
		return &at
	}

	return nil
}
