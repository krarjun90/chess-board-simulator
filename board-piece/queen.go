package board_piece

type Queen struct {
	DefaultPiece
}

func NewQueen() Queen {
	return Queen{
		DefaultPiece: DefaultPiece{
			allowedMoves: []Move{
				{hPos: 0, vPos: 1},
				{hPos: 0, vPos: -1},
				{hPos: 1, vPos: 1},
				{hPos: 1, vPos: 0},
				{hPos: 1, vPos: -1},
				{hPos: -1, vPos: 1},
				{hPos: -1, vPos: 0},
				{hPos: -1, vPos: -1},
			},
			canMoveMultiplePlaces: true,
		}}
}

func (q Queen) GetAllowedMoveDirections() []Move {
	return q.allowedMoves
}

func (q Queen) CanMoveMultiplePlaces() bool {
	return q.canMoveMultiplePlaces
}
