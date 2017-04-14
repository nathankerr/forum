package forum

type User struct {
	Uuid     string `json:"uuid" db:"uuid"`
	Username string `json:"username" db:"username"`
	Password string `json:"-"`
	Created  int    `json:"created" db:"created"`
	Removed  int    `json:"removed" db:"removed"`
}
