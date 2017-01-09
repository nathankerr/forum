package forum

type Thread struct {
	ID      int
	User    int
	Title   string
	Board   int
	Removed int
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
