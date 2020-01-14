package client

import (
	"github.com/TylerGrey/studyhub/internal/mysql/repo"
)

// Client ...
type Client struct {
	UserRepo repo.UserRepository
}
