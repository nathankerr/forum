package main

type User struct {
	ID       int
	Username string
	Password string
	Role     int
	Removed  int
}

type UserService interface {
	User(id int) (*User, error)
	Users() ([]*User, error)
	CreateUser(u *User) error
	DeleteUser(id int) error
}
