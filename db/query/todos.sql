-- name: CreateTodo :one
INSERT INTO todos (name, description, user_id) VALUES ($1, $2, $3) RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: EditTodo :one
UPDATE todos
SET name=$2, description=$3
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id=$1;