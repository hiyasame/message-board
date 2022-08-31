package dao

func InsertUser(user *User) error {
	result := dB.Create(user)
	return result.Error
}

func SelectUserByEmail(user *User, email string) (err error) {
	result := dB.First(user, "email = ?", email)
	err = result.Error
	return
}
