// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createTodoStmt, err = db.PrepareContext(ctx, createTodo); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTodo: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteTodoStmt, err = db.PrepareContext(ctx, deleteTodo); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTodo: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.editTodoStmt, err = db.PrepareContext(ctx, editTodo); err != nil {
		return nil, fmt.Errorf("error preparing query EditTodo: %w", err)
	}
	if q.editUserStmt, err = db.PrepareContext(ctx, editUser); err != nil {
		return nil, fmt.Errorf("error preparing query EditUser: %w", err)
	}
	if q.getTodoStmt, err = db.PrepareContext(ctx, getTodo); err != nil {
		return nil, fmt.Errorf("error preparing query GetTodo: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.listTodosStmt, err = db.PrepareContext(ctx, listTodos); err != nil {
		return nil, fmt.Errorf("error preparing query ListTodos: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createTodoStmt != nil {
		if cerr := q.createTodoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTodoStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteTodoStmt != nil {
		if cerr := q.deleteTodoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTodoStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.editTodoStmt != nil {
		if cerr := q.editTodoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing editTodoStmt: %w", cerr)
		}
	}
	if q.editUserStmt != nil {
		if cerr := q.editUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing editUserStmt: %w", cerr)
		}
	}
	if q.getTodoStmt != nil {
		if cerr := q.getTodoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTodoStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.listTodosStmt != nil {
		if cerr := q.listTodosStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTodosStmt: %w", cerr)
		}
	}
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db             DBTX
	tx             *sql.Tx
	createTodoStmt *sql.Stmt
	createUserStmt *sql.Stmt
	deleteTodoStmt *sql.Stmt
	deleteUserStmt *sql.Stmt
	editTodoStmt   *sql.Stmt
	editUserStmt   *sql.Stmt
	getTodoStmt    *sql.Stmt
	getUserStmt    *sql.Stmt
	listTodosStmt  *sql.Stmt
	listUsersStmt  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:             tx,
		tx:             tx,
		createTodoStmt: q.createTodoStmt,
		createUserStmt: q.createUserStmt,
		deleteTodoStmt: q.deleteTodoStmt,
		deleteUserStmt: q.deleteUserStmt,
		editTodoStmt:   q.editTodoStmt,
		editUserStmt:   q.editUserStmt,
		getTodoStmt:    q.getTodoStmt,
		getUserStmt:    q.getUserStmt,
		listTodosStmt:  q.listTodosStmt,
		listUsersStmt:  q.listUsersStmt,
	}
}
