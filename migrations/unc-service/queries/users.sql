-- name: GetUserByUsernameAndEmail :one
SELECT * FROM users
WHERE username = $1 OR email = $2 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (username, email, password, gender, dob)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: SetIsVerifiedByUsername :one
UPDATE users SET is_verified = $1
WHERE id = $2
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;
