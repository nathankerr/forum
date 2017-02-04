package forum

type User struct {
	Username string `json:"username"`
}

type Users struct {
	Users []User
}
