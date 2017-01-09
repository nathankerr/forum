package forum

type Post struct {
	ID       int
	IsSticky int
	User     int
	Title    string
	Thread   int
	Created  int
	Modified int
	Removed  int
}

type Posts struct {
	Posts []Post
}

type PostService interface {
	Post(id int) (*Post, error)
	Posts() ([]*Post, error)
	CreatePost(u *Post) error
	DeletePost(id int) error
}
