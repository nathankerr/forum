package forum

// Post contains the information of one post.
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

// Posts contains a list of posts.
type Posts struct {
	Posts []Post
}
