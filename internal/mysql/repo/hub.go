package repo

import (
	"fmt"
	"time"

	"github.com/TylerGrey/study-hub/internal/mysql"

	"github.com/jinzhu/gorm"
)

// Hub 사용자
type Hub struct {
	ID         int64 `gorm:"primary_key"`
	Name       string
	Type       string
	CoverImage mysql.JSON
	Images     mysql.JSON
	Tel        string
	Address    string
	Lat        float64
	Lng        float64
	Hours      mysql.JSON
	Cursor     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

// HubRepository Hub 레포지터리 인터페이스
type HubRepository interface {
	Create(hub Hub) (*Hub, error)
	Update(hub Hub) (*Hub, error)
	Delete(id int64) error
	FindByID(id int64) (*Hub, error)

	List(args ListArgs) ([]*Hub, PageInfo, error)
}

// hubRepository 인터페이스 구조체
type hubRepository struct {
	master  *gorm.DB
	replica *gorm.DB
}

// NewHubRepository ...
func NewHubRepository(master *gorm.DB, replica *gorm.DB) HubRepository {
	return &hubRepository{
		master:  master,
		replica: replica,
	}
}

// Create 허브 생성
func (r hubRepository) Create(hub Hub) (*Hub, error) {
	if err := r.master.Table("hub").Create(&hub).Error; err != nil {
		return nil, err
	}

	return &hub, nil
}

// Update 허브 수정
func (r hubRepository) Update(hub Hub) (*Hub, error) {
	if err := r.master.Table("hub").Model(&hub).Update(&hub).Error; err != nil {
		return nil, err
	}

	response := &Hub{}
	if err := r.master.Table("hub").First(&response, hub.ID).Error; err != nil {
		return nil, err
	}

	return response, nil
}

// Delete 허브 수정
func (r hubRepository) Delete(id int64) error {
	if err := r.master.Table("hub").Delete(&Hub{
		ID: id,
	}).Error; err != nil {
		return err
	}

	return nil
}

// FindByID ID로 Hub 조회
func (r hubRepository) FindByID(id int64) (*Hub, error) {
	hub := &Hub{}

	err := r.replica.Table("hub").Where("id = ?", id).Find(hub).Error
	if err != nil {
		return nil, err
	}

	return hub, nil
}

// List Hub List 조회
func (r hubRepository) List(args ListArgs) ([]*Hub, PageInfo, error) {
	var (
		hubs                    []*Hub
		count                   int32
		baseCursor              = getBaseCursor(args)
		cursor, cursorDirection = getCursor(args)
		field, direction        = getOrderBy(args)
		limit                   = getLimit(args)
	)

	tx := r.replica.Table("hub")
	tx = tx.Select(fmt.Sprintf("*, TO_BASE64(%s) as `cursor`", baseCursor))
	tx.Where("deleted_at IS NULL").Count(&count)

	if len(cursor) > 0 {
		tx = tx.Where(fmt.Sprintf("%s %s FROM_BASE64(?)", baseCursor, cursorDirection), cursor)
	}

	tx = tx.Order(fmt.Sprintf("%s %s", field, direction))
	tx = tx.Order(fmt.Sprintf("id %s", direction))
	tx = tx.Limit(limit + 1)

	if err := tx.Find(&hubs).Error; err != nil {
		return nil, PageInfo{}, err
	}

	list, pageInfo := getPageInfo(args, limit, &hubs)
	pageInfo.Total = count

	return list.([]*Hub), pageInfo, nil
}
