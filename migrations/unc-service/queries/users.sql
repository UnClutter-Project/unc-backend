-- name: GetUserByUsernameAndEmail :one
SELECT * FROM users
WHERE username = $1 OR email = $2 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (username, email, password, gender, dob)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users
<<<<<<< HEAD
WHERE username = $1 LIMIT 1;
=======
WHERE username = $1 LIMIT 1;

-- name: SetIsVerifiedByUsername :one
UPDATE users SET is_verified = $1
WHERE username = $2
RETURNING *;
>>>>>>> 33d87161295dd48ec2ee2666b4bc36132b215f81
