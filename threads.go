package forum

type Thread struct {
	ID      int    `json:"id"`
	User    int    `json:"user"`
	Title   string `json:"title"`
	Board   int    `json:"board"`
	Removed int    `json:"-"`
}

type Threads struct {
	Threads []Thread
}

type ThreadService interface {
	Thread(id int) (*Thread, error)
	Threads() ([]*Thread, error)
	CreateThread(u *Thread) error
	DeleteThread(id int) error
}
