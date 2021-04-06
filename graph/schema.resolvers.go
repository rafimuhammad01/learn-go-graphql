package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rafimuhammad01/learn-go-graphql.git/graph/generated"
	"github.com/rafimuhammad01/learn-go-graphql.git/graph/model"
	"github.com/rafimuhammad01/learn-go-graphql.git/graph/repo"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return db.Create(&input), nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, input model.NewTodo) (*model.Todo, error) {
	return db.Update(id, &input), nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input string) (*model.Todo, error) {
	return db.Delete(input), nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return db.FindAll(), nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	return db.FindByID(id), nil
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
var db = repo.NewTodo()
