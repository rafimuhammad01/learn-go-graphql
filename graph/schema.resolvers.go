package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	db "github.com/rafimuhammad01/learn-go-graphql/db/sqlc"
	"github.com/rafimuhammad01/learn-go-graphql/graph/generated"
	"github.com/rafimuhammad01/learn-go-graphql/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*db.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListUser(ctx context.Context) ([]*db.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*db.User, error) {
	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	result, err := r.Queries.GetUser(ctx, n)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) User(ctx context.Context) ([]*db.User, error) {
	panic(fmt.Errorf("not implemented"))
}
