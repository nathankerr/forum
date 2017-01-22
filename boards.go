package forum

type Board struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category int    `json:"category"`
	Position int    `json:"position"`
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
