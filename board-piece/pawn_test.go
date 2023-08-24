package board_piece_test

import (
	p "chess-board-simulator/board-piece"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPawn(t *testing.T) {
	pawn := p.NewPawn()
	assert.NotNil(t, pawn)
	assert.False(t, pawn.CanMoveMultiplePlaces())
	assert.Equal(t, 1, len(pawn.GetAllowedMoveDirections()))
}
