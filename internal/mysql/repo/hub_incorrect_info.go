package repo

import (
	"time"

	"github.com/jinzhu/gorm"
)

// HubIncorrectInfo 잘못된 정보 신고
type HubIncorrectInfo struct {
	ID        uint64 `gorm:"primary_key"`
	HubID     uint64
	UserID    uint64
	Message   string
	IsFixed   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// HubIncorrectInfoRepository Hub 레포지터리 인터페이스
type HubIncorrectInfoRepository interface {
	Create(info HubIncorrectInfo) (*HubIncorrectInfo, error)
}

// hubIncorrectInfoRepository 인터페이스 구조체
type hubIncorrectInfoRepository struct {
	master  *gorm.DB
	replica *gorm.DB
}

// NewHubIncorrectInfoRepository ...
func NewHubIncorrectInfoRepository(master *gorm.DB, replica *gorm.DB) HubIncorrectInfoRepository {
	return &hubIncorrectInfoRepository{
		master:  master,
		replica: replica,
	}
}

// Create ...
func (r hubIncorrectInfoRepository) Create(info HubIncorrectInfo) (*HubIncorrectInfo, error) {
	if err := r.master.Table("hub_incorrect_info").Create(&info).Error; err != nil {
		return nil, err
	}

	return &info, nil
}
