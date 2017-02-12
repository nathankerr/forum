package forum

type Information struct {
	Title string
}

type Users struct {
	Users []string
}

type User struct {
	Username string `json:"username"`
}

type NewUser struct {
	User
	Password string `json:"password"`
}
