package resolver

import (
	"github.com/TylerGrey/studyhub/internal/resolver/model"
)

// HubHour ...
type HubHour struct {
	Data model.HubHour
}

// Label ...
func (h HubHour) Label() string {
	return h.Data.Label
}

// Value ...
func (h HubHour) Value() string {
	return h.Data.Value
}
