package board_piece

type Queen struct {
	DefaultPiece
}

func NewQueen() Piece {
	return &Queen{
		DefaultPiece: DefaultPiece{
			allowedMoves: []Move{
				NewMove(0, 1),
				NewMove(0, -1),
				NewMove(1, 1),
				NewMove(1, 0),
				NewMove(1, -1),
				NewMove(-1, 1),
				NewMove(-1, 0),
				NewMove(-1, -1),
			},
			canMoveMultiplePlaces: true,
		}}
}

func (q *Queen) GetAllowedMoveDirections() []Move {
	return q.allowedMoves
}

func (q *Queen) CanMoveMultiplePlaces() bool {
	return q.canMoveMultiplePlaces
}
