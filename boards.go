package forum

// Board contains the information of one board.
type Board struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category int    `json:"category"`
	Position int    `json:"position"`
}

// Boards contains a list of boards.
type Boards struct {
	Boards []Board
}
