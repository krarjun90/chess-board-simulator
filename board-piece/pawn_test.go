package board_piece_test

import (
	"github.com/stretchr/testify/assert"
	p "technogise/chess-board/board-piece"
	"testing"
)

func TestNewPawn(t *testing.T) {
	pawn := p.NewPawn()
	assert.NotNil(t, pawn)
	assert.False(t, pawn.CanMoveMultiplePlaces())
	assert.Equal(t, 1, len(pawn.GetAllowedMoveDirections()))
}
