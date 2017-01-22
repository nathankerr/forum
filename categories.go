package forum

type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
}

type Categories struct {
	Categories []Category
}

type CategoryService interface {
	Category(id int) (*Category, error)
	Categories() ([]*Category, error)
	CreateCategory(u *Category) error
	DeleteCategory(id int) error
}
