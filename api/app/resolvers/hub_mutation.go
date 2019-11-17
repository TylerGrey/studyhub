package resolvers

import (
	"encoding/json"
	"strconv"

	"github.com/TylerGrey/study-hub/api/app/resolvers/args"
	"github.com/TylerGrey/study-hub/internal/mysql/repo"
	"github.com/TylerGrey/study-hub/internal/resolver"
)

// CreateHub Hub 생성
func (r *Resolver) CreateHub(input args.CreateHubInput) (*resolver.Hub, error) {
	m := repo.Hub{
		Name:    input.Input.Name,
		Type:    input.Input.Type,
		Address: input.Input.Address.Address,
		Lat:     input.Input.Address.Lat,
		Lng:     input.Input.Address.Lng,
	}

	// TODO: 이미지 처리
	coverImage, _ := json.Marshal(input.Input.CoverImage)
	m.CoverImage = coverImage

	if input.Input.Tel != nil {
		m.Tel = input.Input.Tel
	}

	if input.Input.Images != nil {
		// TODO: 이미지 처리
	}

	if input.Input.Hours != nil {
		hours, _ := json.Marshal(*input.Input.Hours)
		m.Hours = hours
	}

	hub, err := r.HubRepo.Create(m)
	if err != nil {
		return nil, err
	}

	return &resolver.Hub{
		Data: *hub,
	}, nil
}

// UpdateHub Hub 수정
func (r *Resolver) UpdateHub(input args.UpdateHubInput) (*resolver.Hub, error) {
	id, err := strconv.ParseUint(input.Input.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	m := repo.Hub{
		ID: id,
	}

	if input.Input.Name != nil {
		m.Name = *input.Input.Name
	}

	if input.Input.Type != nil {
		m.Type = *input.Input.Type
	}

	if input.Input.CoverImage != nil {
		// TODO: 이미지 처리
	}

	if input.Input.Images != nil {
		// TODO: 이미지 처리
	}

	if input.Input.Tel != nil {
		m.Tel = input.Input.Tel
	}

	if input.Input.Address != nil {
		m.Address = input.Input.Address.Address
		m.Lat = input.Input.Address.Lat
		m.Lng = input.Input.Address.Lng
	}

	if input.Input.Hours != nil {
		hours, _ := json.Marshal(*input.Input.Hours)
		m.Hours = hours
	}

	hub, err := r.HubRepo.Update(m)
	if err != nil {
		return nil, err
	}

	return &resolver.Hub{
		Data: *hub,
	}, nil
}

// DeleteHub Hub 삭제
func (r *Resolver) DeleteHub(args struct {
	ID string
}) (bool, error) {
	id, err := strconv.ParseUint(args.ID, 10, 64)
	if err != nil {
		return false, err
	}

	if err = r.HubRepo.Delete(id); err != nil {
		return false, err
	}

	return true, nil
}

// AddHubIncorrectInfo 잘못된 정보 수정 요청
func (r *Resolver) AddHubIncorrectInfo(input args.AddHubIncorrectInfoInput) (*resolver.HubIncorrectInfo, error) {
	hubID, err := strconv.ParseUint(input.Input.HubID, 10, 64)
	if err != nil {
		return nil, err
	}

	// TODO: 로그인 적용 하면, 토큰으로 User id 받아오도록 수정 필요
	m := repo.HubIncorrectInfo{
		HubID:   hubID,
		UserID:  uint64(1),
		Message: input.Input.Message,
	}

	info, err := r.HubIncorrectInfoRepo.Create(m)
	if err != nil {
		return nil, err
	}

	return &resolver.HubIncorrectInfo{
		Data: *info,
	}, nil
}
