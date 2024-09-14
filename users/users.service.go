package users

import (
	"fmt"

	"github.com/9thDuck/ecommerce-api.git/common"
	"github.com/google/uuid"
)

func createUser(user *User) error {
	err := user.hashPassword()
	if err != nil {
		return err
	}
	// returns an error / nil
	return user.Create()
}

func loginUser(user *User) (accessToken string, refreshToken string, err error) {
	if err := user.getUser(); err != nil {
		return "", "", err
	}

	if err := user.verifyLoginCredentials(); err != nil {
		return "", "", err
	}

	if user.Banned {
		return "", "", fmt.Errorf("you have been banned from the platform")
	}

	accessToken, refreshToken, err = user.GenerateToken()
	if err != nil {
		common.LogCustomError(fmt.Sprintf("error generating tokens for userId %s\n", user.ID), err)
		return "", "", fmt.Errorf("error loggin you in")

	}

	session := Session{
		UserID:       user.ID,
		RefreshToken: refreshToken,
	}

	if err := session.Create(); err != nil {
		fmt.Println("error creating session")
		return "", "", fmt.Errorf("error loggin you in")
	}

	return accessToken, refreshToken, err
}

func getUserDetails(user *User) error {
	if err := user.getUser(); err != nil {
		return err
	}
	return nil
}

func logoutUser(refreshToken string) error {
	session := Session{
		RefreshToken: refreshToken,
	}
	return session.End()
}

func logoutOfAllSessions(userId uuid.UUID) error {
	return endAll(userId)
}

func banUser(userId uuid.UUID) error {
	return banIfNotAdmin(userId)
}
