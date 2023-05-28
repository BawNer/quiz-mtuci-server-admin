package usecase

import (
	"fmt"
	"quiz-mtuci-server/internal/entity"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	key []byte
}

func NewJWT(key []byte) JWT {
	return JWT{
		key: key,
	}
}

func (j *JWT) Create(ttl time.Duration, content interface{}) (string, error) {
	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["dat"] = content
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.key)
	if err != nil {
		return "", fmt.Errorf("create %w", err)
	}

	return token, nil
}

func (j *JWT) Validate(token string) (map[string]interface{}, error) {
	t, e := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpocted method %s", jwtToken.Header["alg"])
		}

		return j.key, nil
	})
	if e != nil {
		return nil, fmt.Errorf("validate: %w", e)
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}

	return claims["dat"].(map[string]interface{}), nil
}

func (j *JWT) Parse(payload map[string]interface{}) entity.User {
	id := payload["id"].(float64)
	email := payload["email"].(string)
	name := payload["name"].(string)
	zach := payload["numberZach"].(string)
	is_student := payload["isStudent"].(float64)
	user := entity.User{
		ID:         int(id),
		Email:      email,
		Name:       name,
		NumberZach: zach,
		IsStudent:  int(is_student),
	}

	return user
}
