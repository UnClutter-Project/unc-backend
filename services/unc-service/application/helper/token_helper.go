package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat":     time.Now().Unix(),                // Issued at
	}
	t, err := generateJwt(claims)
	if err != nil {
		return "", err
	}

	return t, nil
}

func GenerateRefreshToken(id string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(), // Expiration time
		"iat":     time.Now().Unix(),                          // Issued at
	}
	t, err := generateJwt(claims)
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyAndParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func generateJwt(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetUserFromJwt(token *jwt.Token) string {
	claims := token.Claims.(jwt.MapClaims)
	user_id, ok := claims["user_id"].(string) // Ensure the key matches your token generation
	if !ok {
		return ""
	}

	return user_id
}
