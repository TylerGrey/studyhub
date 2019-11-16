package resolver

import (
	"encoding/json"
	"strconv"

	"github.com/TylerGrey/study-hub/internal/mysql/repo"
	"github.com/TylerGrey/study-hub/internal/resolver/model"
	"github.com/graph-gophers/graphql-go"
)

// Hub ...
type Hub struct {
	Data repo.Hub
}

// ID ...
func (h Hub) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(h.Data.ID)))
}

// Name ...
func (h Hub) Name() string {
	return h.Data.Name
}

// CoverImage ...
func (h Hub) CoverImage() Image {
	var image model.Image
	json.Unmarshal(h.Data.CoverImage, &image)
	return Image{
		Data: image,
	}
}

// Images ...
func (h Hub) Images() *[]Image {
	resolvers := []Image{}

	var images []model.Image
	json.Unmarshal(h.Data.Images, &images)

	for _, image := range images {
		resolvers = append(resolvers, Image{
			Data: image,
		})
	}

	return &resolvers
}

// Tel ...
func (h Hub) Tel() *string {
	return &h.Data.Tel
}

// Address ...
func (h Hub) Address() Address {
	return Address{
		Data: model.Address{
			Address: h.Data.Address,
			Lat:     h.Data.Lat,
			Lng:     h.Data.Lng,
		},
	}
}

// Hours ...
func (h Hub) Hours() *[]HubHour {
	resolvers := []HubHour{}

	var hours []model.HubHour
	json.Unmarshal(h.Data.Hours, &hours)

	for _, hour := range hours {
		resolvers = append(resolvers, HubHour{
			Data: hour,
		})
	}

	return &resolvers
}

// CreatedAt ...
func (h Hub) CreatedAt() string {
	return h.Data.CreatedAt.String()
}

// UpdatedAt ...
func (h Hub) UpdatedAt() string {
	return h.Data.UpdatedAt.String()
}

// DeletedAt ...
func (h Hub) DeletedAt() *string {
	if h.Data.DeletedAt != nil {
		at := h.Data.DeletedAt.String()
		return &at
	}

	return nil
}
