package forum

type Board struct {
	ID       int
	Name     string
	Category int
	Position int
}

type Boards struct {
	Boards []Board
}

type BoardService interface {
	Board(id int) (*Board, error)
	Boards() ([]*Board, error)
	CreateBoard(u *Board) error
	DeleteBoard(id int) error
}
