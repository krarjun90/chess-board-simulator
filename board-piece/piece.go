package board_piece

import (
	"errors"
	"strings"
)

type Move struct {
	hPos int
	vPos int
}

func NewMove(hPos, vPos int) Move {
	return Move{hPos: hPos, vPos: vPos}
}

func (m Move) GetHPos() int {
	return m.hPos
}

func (m Move) GetVPos() int {
	return m.vPos
}

type Piece interface {
	GetAllowedMoveDirections() []Move
	CanMoveMultiplePlaces() bool
}

type DefaultPiece struct {
	allowedMoves          []Move
	canMoveMultiplePlaces bool
}

func ParsePiece(p string) (Piece, error) {
	switch strings.ToLower(p) {
	case "pawn":
		return NewPawn(), nil
	case "king":
		return NewKing(), nil
	case "queen":
		return NewQueen(), nil
	case "knight":
		return NewKnight(), nil
	default:
		return nil, errors.New("piece not supported")
	}
}
