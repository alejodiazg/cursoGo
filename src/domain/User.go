package domain

type User struct {
	Username string
	Name     string
	Email    string
	Password string
}

func NewUser(username, name, email, password string) *User{
	return &User{username, name, email, password}
}