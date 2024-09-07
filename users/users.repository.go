package users

import (
	"fmt"

	"github.com/9thDuck/ecommerce-api.git/db"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

func (user *User) Create() error {
	err := db.Instance.Create(user).Error
	if err != nil {
		if pgError, isPgError := err.(*pgconn.PgError); isPgError {
			fmt.Println(pgError.ConstraintName)
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
	user.Password = string(hashedPassword)
	return nil
}

func New(username, email, password string) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
	}
}
