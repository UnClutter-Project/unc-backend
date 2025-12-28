-- name: GetValidOTPByUsernameAndCode :one
SELECT * FROM otp
WHERE user_id = $1 AND code = $2 AND expired_at > NOW() LIMIT 1;

-- name: CreateOTP :one
INSERT INTO otp (user_id, code)
VALUES ($1, $2)
RETURNING *;