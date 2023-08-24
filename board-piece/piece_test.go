package board_piece_test

import (
	"github.com/stretchr/testify/assert"
	p "technogise/chess-board/board-piece"
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
