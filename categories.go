package forum

// Category contains the information of one category.
type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
}

// Categories contains a list of categories.
type Categories struct {
	Categories []Category
}
