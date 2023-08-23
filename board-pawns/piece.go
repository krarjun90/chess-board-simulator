package board_pawns

import (
	"errors"
	"strings"
)

type Piece int

const (
	Pawn Piece = iota
	King
	Queen
)

func GetPiece(p string) (Piece, error) {
	switch strings.ToLower(p) {
	case "pawn":
		return Pawn, nil
	case "king":
		return King, nil
	case "queen":
		return Queen, nil
	default:
		return Piece(-1), errors.New("piece not supported")
	}
}
