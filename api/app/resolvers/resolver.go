package resolvers

import (
	"github.com/TylerGrey/hub-api/api/app/repo/mysql"
)

// Resolver ...
type Resolver struct {
	UserRepo mysql.UserRepository
}
