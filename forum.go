package forum

type User struct {
	Uuid     string `json:"uuid" db:"uuid"`
	Username string `json:"username" db:"username"`
	Password string `json:"-"`
	Created  int    `json:"created" db:"created"`
	Removed  int    `json:"removed" db:"removed"`
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error *Error      `json:"error,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
