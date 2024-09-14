package middleware

import (
	"github.com/9thDuck/ecommerce-api.git/db"
	"github.com/9thDuck/ecommerce-api.git/entities"
	"github.com/google/uuid"
)

type UserWithSession struct {
	User    entities.User
	Session entities.Session
}

func getUserSession(refreshToken string, userId uuid.UUID) (*UserWithSession, error) {
	session := entities.Session{RefreshToken: refreshToken, UserID: userId}
	err := db.Instance.Where(&session).Preload("User").First(&session).Error
	if err != nil {
		return nil, err
	}
	return &UserWithSession{User: session.User, Session: session}, nil
}
