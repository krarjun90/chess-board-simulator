package board_piece

type Pawn struct {
	DefaultPiece
}

func NewPawn() Pawn {
	return Pawn{
		DefaultPiece: DefaultPiece{
			allowedMoves: []Move{
				{hPos: 0, vPos: 1},
			},
			canMoveMultiplePlaces: false,
		},
	}
}

func (p Pawn) GetAllowedMoveDirections() []Move {
	return p.allowedMoves
}

func (p Pawn) CanMoveMultiplePlaces() bool {
	return p.canMoveMultiplePlaces
}
