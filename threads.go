package forum

// Thread contains the information of one thread.
type Thread struct {
	ID      int    `json:"id"`
	User    int    `json:"user"`
	Title   string `json:"title"`
	Board   int    `json:"board"`
	Removed int    `json:"-"`
}

// Threads contains a list of threads.
type Threads struct {
	Threads []Thread
}
