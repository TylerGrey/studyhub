package resolvers

import (
	"context"
)

type ByeInput struct {
	Input WordArgs
}

type WordArgs struct {
	Word string
}

// Hello Hello query
func (r *Resolver) Hello(ctx context.Context) HelloResolver {
	return HelloResolver{
		word: "World",
	}
}

// Bye Bye Mutation
func (r *Resolver) Bye(ctx context.Context, args ByeInput) HelloResolver {
	return HelloResolver{
		word: args.Input.Word,
	}
}
