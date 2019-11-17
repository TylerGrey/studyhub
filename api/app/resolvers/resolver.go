package resolvers

import (
	"github.com/TylerGrey/study-hub/internal/mysql/repo"
)

// Resolver ...
type Resolver struct {
	UserRepo             repo.UserRepository
	HubRepo              repo.HubRepository
	HubIncorrectInfoRepo repo.HubIncorrectInfoRepository
}
