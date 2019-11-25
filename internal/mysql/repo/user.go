package repo

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User 사용자
type User struct {
	ID        uint64 `gorm:"primary_key"`
	UUID      string
	Email     string
	Password  string
	FirstName *string
	LastName  *string
	Nickname  string
	Mobile    string
	Birth     *string
	Gender    *string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// UserRepository User 레포지터리 인터페이스
type UserRepository interface {
	Create(user User) (*User, error)
	List(page int32, limit int32) ([]*User, error)
	FindByID(id uint64) (*User, error)
}

// userRepository 인터페이스 구조체
type userRepository struct {
	master  *gorm.DB
	replica *gorm.DB
}

// NewUserRepository ...
func NewUserRepository(master *gorm.DB, replica *gorm.DB) UserRepository {
	return &userRepository{
		master:  master,
		replica: replica,
	}
}

// Create 유저 생성
func (r userRepository) Create(user User) (*User, error) {
	err := r.master.Table("user").Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// List 유저 리스트 조회
func (r userRepository) List(page int32, limit int32) ([]*User, error) {
	users := []*User{}

	err := r.replica.Table("user").Scan(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

// FindByID ID로 유저 조회
func (r userRepository) FindByID(id uint64) (*User, error) {
	user := &User{}

	err := r.replica.
		Table("user").
		Where("id = ?", id).
		Find(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
