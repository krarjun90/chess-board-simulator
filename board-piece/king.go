package board_piece

type King struct {
	DefaultPiece
}

func NewKing() Piece {
	return &King{
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
			canMoveMultiplePlaces: false,
		}}
}

func (k *King) GetAllowedMoveDirections() []Move {
	return k.allowedMoves
}

func (k *King) CanMoveMultiplePlaces() bool {
	return k.canMoveMultiplePlaces
}
