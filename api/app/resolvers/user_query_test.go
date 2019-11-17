package resolvers_test

import (
	"testing"
)

func TestUser(t *testing.T) {
	_, err := resolver.User(struct{ ID string }{
		ID: "1",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
