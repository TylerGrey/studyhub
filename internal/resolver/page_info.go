package resolver

import (
	"github.com/TylerGrey/study-hub/internal/resolver/model"
)

// PageInfo 공통 페이지 정보
type PageInfo struct {
	Data model.PageInfo
}

// StartCursor ...
func (p PageInfo) StartCursor() *string {
	return p.Data.StartCursor
}

// EndCursor ...
func (p PageInfo) EndCursor() *string {
	return p.Data.EndCursor
}

// HasNextPage ...
func (p PageInfo) HasNextPage() bool {
	return p.Data.HasNextPage
}

// HasPreviousPage ...
func (p PageInfo) HasPreviousPage() bool {
	return p.Data.HasPreviousPage
}
