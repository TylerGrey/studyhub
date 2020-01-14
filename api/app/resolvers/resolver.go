package resolvers

import (
	"github.com/TylerGrey/studyhub/internal/mysql/repo"
)

// Resolver ...
type Resolver struct {
	UserRepo             repo.UserRepository
	HubRepo              repo.HubRepository
	HubIncorrectInfoRepo repo.HubIncorrectInfoRepository
	HubReviewRepo        repo.HubReviewRepository
}
