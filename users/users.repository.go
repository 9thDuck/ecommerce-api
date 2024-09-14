package users

import (
	"errors"

	"github.com/9thDuck/ecommerce-api.git/auth"
	"github.com/9thDuck/ecommerce-api.git/db"
	"github.com/9thDuck/ecommerce-api.git/entities"
	"github.com/9thDuck/ecommerce-api.git/middleware"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type User entities.User
type Session entities.Session

func (user User) Create() error {
	err := db.Instance.Create(&user).Error
	if err != nil {
		if pgError, isPgError := err.(*pgconn.PgError); isPgError {
			err = db.TranslatePgErrors(pgError)
		}

		return err
	}
	return nil
}

// We are hashing the password before creating the User
func (user *User) hashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.HashedPassword = string(hashedPassword)
	return nil
}

func New(username, email, password string, role int) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}

}

func (user *User) getUser() error {
	if foundUsers := db.Instance.Limit(1).Find(&user, &user).RowsAffected; foundUsers == 0 {
		return errors.New("User not found")
	}
	return nil
}

func (user *User) verifyLoginCredentials() error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(user.Password)); err != nil {
		return err
	}
	return nil
}

func (user *User) GenerateToken() (accessTokenStr string, refreshTokenStr string, err error) {
	claims := auth.TokenClaims{
		ID:   user.ID,
		Role: user.Role,
	}
	return auth.GenerateToken(claims)
}

func (session *Session) Create() error {
	err := db.Instance.Create(&session).Error
	if err != nil {
		if pgError, isPgError := err.(*pgconn.PgError); isPgError {
			err = db.TranslatePgErrors(pgError)
		}

		return err
	}
	return nil
}

func (session *Session) End() error {
	return db.Instance.Model(&Session{}).Where(&session).Update("logged_out", true).Error
}

func endAll(userId uuid.UUID) error {
	return db.Instance.Model(&Session{}).Where("user_id = ? AND logged_out = ?", userId, false).Update("logged_out", true).Error
}

func banIfNotAdmin(userId uuid.UUID) error {
	return db.Instance.Model(&User{}).Where("id = ? AND role != ?", userId, middleware.ADMIN).Update("banned", true).Error
}
