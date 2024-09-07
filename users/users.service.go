package users

func createUser(user *User) error {
	err := user.hashPassword()
	if err != nil {
		return err
	}
	// returns an error / nil
	return user.Create()
}
