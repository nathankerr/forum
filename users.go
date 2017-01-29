package forum

// User contains the information of one user.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Role     int    `json:"-"`
	Removed  int    `json:"-"`
}

// Users contains a list of all users.
type Users struct {
	Users []User
}
