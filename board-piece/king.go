package board_piece

type King struct {
	DefaultPiece
}

func NewKing() King {
	return King{
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
			canMoveMultiplePlaces: false,
		}}
}

func (k King) GetAllowedMoveDirections() []Move {
	return k.allowedMoves
}

func (k King) CanMoveMultiplePlaces() bool {
	return k.canMoveMultiplePlaces
}
