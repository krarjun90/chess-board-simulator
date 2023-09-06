package board_piece

type Pawn struct {
	DefaultPiece
}

func NewPawn() Piece {
	return &Pawn{
		DefaultPiece: DefaultPiece{
			allowedMoves: []Move{
				NewMove(0, 1),
			},
			canMoveMultiplePlaces: false,
		},
	}
}

func (p *Pawn) GetAllowedMoveDirections() []Move {
	return p.allowedMoves
}

func (p *Pawn) CanMoveMultiplePlaces() bool {
	return p.canMoveMultiplePlaces
}
