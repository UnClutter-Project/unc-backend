-- name: GetValidOTPByUsernameAndToken :one
SELECT * FROM otp
WHERE user_id = $1 AND token = $2 AND expired_at > NOW() LIMIT 1;