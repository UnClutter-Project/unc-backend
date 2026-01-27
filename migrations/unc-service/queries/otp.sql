-- name: GetValidOTPByCode :one
SELECT * FROM otp
WHERE code = $1 AND expired_at > NOW() LIMIT 1;

-- name: CreateOTP :one
INSERT INTO otp (user_id, code)
VALUES ($1, $2)
RETURNING *;