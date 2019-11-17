package resolvers_test

import (
	"github.com/TylerGrey/study-hub/api/app/resolvers/args"
	"testing"
)

func TestCreateUser(t *testing.T) {
	_, err := resolver.CreateUser(args.CreateUserInput{
		Input: args.CreateUserArgs{
			Email:    "tyler.grey76@gmail.com",
			Name:     "Tyler Grey",
			Password: "1234qwer",
		},
	})
	if err != nil {
		t.Error(err.Error())
	}
}
