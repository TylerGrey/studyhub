package resolver

import (
	"github.com/TylerGrey/study-hub/internal/mysql/repo"
	"github.com/TylerGrey/study-hub/internal/resolver/model"
)

// HubConnection ...
type HubConnection struct {
	Data []*repo.Hub
	Page repo.PageInfo
}

// Edges ...
func (h HubConnection) Edges() *[]*HubEdge {
	edges := []*HubEdge{}

	for _, d := range h.Data {
		edges = append(edges, &HubEdge{
			Data: *d,
		})
	}

	return &edges
}

// Nodes ...
func (h HubConnection) Nodes() *[]*Hub {
	nodes := []*Hub{}

	for _, d := range h.Data {
		nodes = append(nodes, &Hub{
			Data: *d,
		})
	}

	return &nodes
}

// PageInfo ...
func (h HubConnection) PageInfo() PageInfo {
	data := model.PageInfo{
		HasNextPage:     h.Page.HasNext,
		HasPreviousPage: h.Page.HasPrev,
	}

	if len(h.Data) > 0 {
		data.StartCursor = &h.Data[0].Cursor
		data.EndCursor = &h.Data[len(h.Data)-1].Cursor
	}

	return PageInfo{
		Data: data,
	}
}

// TotalCount ...
func (h HubConnection) TotalCount() int32 {
	return h.Page.Total
}
