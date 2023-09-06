package board_piece

type Knight struct {
	DefaultPiece
}

func NewKnight() Piece {
	return &Knight{
		DefaultPiece: DefaultPiece{
			allowedMoves: []Move{
				NewMove(2, 1),
				NewMove(2, -1),
				NewMove(-2, 1),
				NewMove(-2, -1),
				NewMove(1, 2),
				NewMove(1, -2),
				NewMove(-1, 2),
				NewMove(-1, -2),
			},
			canMoveMultiplePlaces: false,
		}}
}

func (k *Knight) GetAllowedMoveDirections() []Move {
	return k.allowedMoves
}

func (k *Knight) CanMoveMultiplePlaces() bool {
	return k.canMoveMultiplePlaces
}
