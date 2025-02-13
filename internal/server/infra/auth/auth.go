package auth

import (
	"GophKeeper/internal/server/infra/config"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	secret []byte
}

var ErrInvalidToken = errors.New("invalid token")

func NewService(c config.Config) *Service {
	return &Service{secret: []byte(c.JWTSecret)}
}

func (s *Service) IssueToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iss": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}, nil)

	signed, err := token.SignedString(s.secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}
	return signed, nil
}

func (s *Service) ValidateToken(token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secret, nil
	})
	if err != nil {
		return "", errors.Join(ErrInvalidToken, err)
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		id, ok := claims["sub"].(string)
		if !ok {
			return "", fmt.Errorf("%w: failed to extract id from claims", ErrInvalidToken)
		}

		return id, nil
	}

	return "", ErrInvalidToken
}
