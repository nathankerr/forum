package forum

type Overview struct {
	Users []string
}

type User struct {
	Username string `json:"username"`
}
