package resolver

import (
	"github.com/TylerGrey/study-hub/internal/mysql/repo"
)

// HubEdge ...
type HubEdge struct {
	Data repo.Hub
}

// Cursor ...
func (h HubEdge) Cursor() string {
	// TODO: encrypt MD5
	return h.Data.Cursor
}

// Node ...
func (h HubEdge) Node() *Hub {
	return &Hub{
		Data: h.Data,
	}
}
