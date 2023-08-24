package board_piece_test

import (
	p "chess-board-simulator/board-piece"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueen(t *testing.T) {
	q := p.NewQueen()
	assert.NotNil(t, q)
	assert.True(t, q.CanMoveMultiplePlaces())
	assert.Equal(t, 8, len(q.GetAllowedMoveDirections()))
}
