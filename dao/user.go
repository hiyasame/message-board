package dao

func InsertUser(user *User) error {
	result := dB.Create(user)
	return result.Error
}
