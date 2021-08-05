package graph

import (
	"context"
	db "github.com/rafimuhammad01/learn-go-graphql/db/sqlc"
	"strconv"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input db.CreateUserParams) (*db.User, error) {
	user, err := r.Queries.CreateUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *mutationResolver) EditUser(ctx context.Context, input db.EditUserParams) (*db.User, error) {
	user, err := r.Queries.EditUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (string, error) {
	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return "", err
	}
	err = r.Queries.DeleteUser(ctx, n)
	if err != nil {
		return "", err
	}

	return "Delete Success", nil
}

func (r *queryResolver) ListUser(ctx context.Context, pageID int, pageSize int) ([]*db.User, error) {
	users, err := r.Queries.ListUsers(ctx, db.ListUsersParams{
		Limit:  int32(pageSize),
		Offset: (int32(pageID) - 1) * int32(pageSize),
	})
	if err != nil {
		return nil, err
	}

	var usersParsing []*db.User
	for i := 0; i < len(users); i++ {
		usersParsing = append(usersParsing, &users[i])
	}

	return usersParsing, nil
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
