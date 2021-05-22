package database

type User struct {
	ID                uint64 `json:"id"`
	Username          string `json:"username"`
	EncryptedPassword string `json:"encrypted_password"`
}

var user User = User{
	ID:                1,
	Username:          "fain",
	EncryptedPassword: "123456",
}

func (u User) SelectOne(username string) *User {
	if username != user.Username {
		return nil
	}
	return &user
}

func (u User) Check(password string) bool {
	return user.EncryptedPassword == password
}
