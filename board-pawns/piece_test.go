package board_pawns_test

import (
	"github.com/stretchr/testify/assert"
	"technogise/chess-board/board-pawns"
	"testing"
)

func TestGetInvalidPiece(t *testing.T) {
	p, err := board_pawns.GetPiece("bishop")
	assert.Equal(t, board_pawns.Piece(-1), p)
	assert.NotNil(t, err)
}

func TestGetDefinedPiece(t *testing.T) {
	var testData = map[string]board_pawns.Piece{
		"KING":  board_pawns.King,
		"QUEEN": board_pawns.Queen,
		"PAWN":  board_pawns.Pawn,
	}

	for k, v := range testData {
		p, err := board_pawns.GetPiece(k)
		assert.Equal(t, v, p)
		assert.Nil(t, err)
	}
}
