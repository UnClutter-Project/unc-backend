-- name: GetUserByUsernameAndEmail :one
SELECT * FROM users
WHERE username = $1 OR email = $2 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (username, email, password, gender, dob)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
