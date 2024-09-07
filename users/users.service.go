package users

import (
	"github.com/9thDuck/ecommerce-api.git/utils"
)

func (user *User) CreateUser() error {
	err := user.hashPassword()
	if err != nil {
		utils.LogCustomError("failed to hash the password", err)
		return err
	}
	// returns an error / nil
	return user.Create()
}
