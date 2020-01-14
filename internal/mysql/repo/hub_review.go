package repo

import (
	"time"

	"github.com/TylerGrey/studyhub/internal/mysql"
	"github.com/jinzhu/gorm"
)

// HubReview Hub 리뷰
type HubReview struct {
	ID        uint64 `gorm:"primary_key"`
	HubID     uint64
	UserID    uint64
	Rating    int32
	Review    string
	Images    mysql.JSON
	Cursor    string `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// HubReviewRepository Hub 레포지터리 인터페이스
type HubReviewRepository interface {
	List(hubID uint64) ([]*HubReview, error)
}

// hubReviewRepository 인터페이스 구조체
type hubReviewRepository struct {
	master  *gorm.DB
	replica *gorm.DB
}

// NewHubReviewRepository ...
func NewHubReviewRepository(master *gorm.DB, replica *gorm.DB) HubReviewRepository {
	return &hubReviewRepository{
		master:  master,
		replica: replica,
	}
}

// List Hub List 조회
func (r hubReviewRepository) List(hubID uint64) ([]*HubReview, error) {
	var hubs []*HubReview
	if err := r.replica.Table("hub_review").Find(&hubs).Error; err != nil {
		return nil, err
	}

	return hubs, nil
}
