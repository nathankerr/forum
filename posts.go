package forum

type Post struct {
	ID       int    `json:"id"`
	IsSticky int    `json:"sticky"`
	User     int    `json:"user"`
	Title    string `json:"title"`
	Thread   int    `json:"thread"`
	Content  string `json:"content"`
	Created  int    `json:"created"`
	Modified int    `json:"modified"`
	Removed  int    `json:"-"`
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
