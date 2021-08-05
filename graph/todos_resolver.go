package graph

import (
	"context"
	db "github.com/rafimuhammad01/learn-go-graphql/db/sqlc"
	"strconv"
)

func (r *mutationResolver) EditTodo(ctx context.Context, input db.EditTodoParams) (*db.Todo, error) {
	todo, err := r.Queries.EditTodo(ctx, input)
	if err != nil {
		return nil, err
	}

	return &todo, err
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (string, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	err = r.Queries.DeleteTodo(ctx, int64(idInt))
	if err != nil {
		return "", err
	}

	return "Delete Success", nil

}

func (r *mutationResolver) CreateTodo(ctx context.Context, input db.CreateTodoParams) (*db.Todo, error) {
	todo, err := r.Queries.CreateTodo(ctx, input)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *queryResolver) ListTodo(ctx context.Context, pageID int, pageSize int) ([]*db.Todo, error) {
	todos, err := r.Queries.ListTodos(ctx, db.ListTodosParams{
		Limit:  int32(pageSize),
		Offset: (int32(pageID) - 1) * int32(pageSize),
	})
	if err != nil {
		return nil, err
	}

	var todosParsing []*db.Todo
	for i := 0; i < len(todos); i++ {
		todosParsing = append(todosParsing, &todos[i])
	}

	return todosParsing, nil
}

func (r *queryResolver) GetTodo(ctx context.Context, id string) (*db.Todo, error) {
	ID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	todo, err := r.Queries.GetTodo(ctx, int64(ID))
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *todoResolver) User(ctx context.Context, obj *db.Todo) (*db.User, error) {
	user, err := r.Queries.GetUser(ctx, obj.UserID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

