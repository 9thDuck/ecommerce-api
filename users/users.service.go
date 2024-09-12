package users

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

	accessToken, refreshToken, err = user.GenerateToken()
	if err != nil {
		return "", "", err

	}

	return accessToken, refreshToken, err
}

func getUserDetails(user *User) error {
	if err := user.getUser(); err != nil {
		return err
	}
	return nil
}
