package resolver

import (
	"github.com/TylerGrey/study-hub/internal/resolver/model"
)

// Address ...
type Address struct {
	Data model.Address
}

// Address ...
func (h Address) Address() string {
	return h.Data.Address
}

// Lat ...
func (h Address) Lat() float64 {
	return h.Data.Lat
}

// Lng ...
func (h Address) Lng() float64 {
	return h.Data.Lng
}
