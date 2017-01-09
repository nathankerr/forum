package forum

type Category struct {
	ID       int
	Name     string
	Position int
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
