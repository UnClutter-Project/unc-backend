-- name: GetUserByUsernameAndEmail :one
SELECT * FROM users
WHERE username = $1 OR email = $2 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (username, email, password)
VALUES ($1, $2, $3)
RETURNING *;
