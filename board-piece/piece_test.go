package board_piece_test

import (
	p "chess-board-simulator/board-piece"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInvalidPiece(t *testing.T) {
	piece, err := p.ParsePiece("bishop")
	assert.Equal(t, nil, piece)
	assert.NotNil(t, err)
}

func TestGetDefinedPiece(t *testing.T) {
	var testData = map[string]p.Piece{
		"KING":  p.NewKing(),
		"QUEEN": p.NewQueen(),
		"PAWN":  p.NewPawn(),
	}

	for k, v := range testData {
		piece, err := p.ParsePiece(k)
		assert.Equal(t, v, piece)
		assert.Nil(t, err)
	}
}
