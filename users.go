package forum

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Role     int    `json:"-"`
	Removed  int    `json:"-"`
}

type Users struct {
	Users []User
}

type UsersResource interface {
	Get() (int, interface{}, error)
	Post() (int, interface{}, error)
	Put() (int, interface{}, error)
	Delete() (int, interface{}, error)
}
