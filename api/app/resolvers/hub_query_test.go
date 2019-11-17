package resolvers_test

import (
	"testing"
)

func TestHub(t *testing.T) {
	_, err := resolver.Hub(struct{ ID string }{
		ID: "25",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
