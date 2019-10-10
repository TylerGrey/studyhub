package resolvers

import "fmt"

type HelloResolver struct {
	word string
}

func (m HelloResolver) Word() string {
	return fmt.Sprintf("Hello! %s", m.word)
}
