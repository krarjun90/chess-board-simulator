package board_piece

type Knight struct {
	DefaultPiece
}

func NewKnight() Knight {
	return Knight{
		DefaultPiece: DefaultPiece{
			allowedMoves: []Move{
				{hPos: 2, vPos: 1},
				{hPos: 2, vPos: -1},
				{hPos: -2, vPos: 1},
				{hPos: -2, vPos: -1},
				{hPos: 1, vPos: 2},
				{hPos: 1, vPos: -2},
				{hPos: -1, vPos: 2},
				{hPos: -1, vPos: -2},
			},
			canMoveMultiplePlaces: false,
		}}
}

func (k Knight) GetAllowedMoveDirections() []Move {
	return k.allowedMoves
}

func (k Knight) CanMoveMultiplePlaces() bool {
	return k.canMoveMultiplePlaces
}
