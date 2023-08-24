package board_piece_test

import (
	p "chess-board-simulator/board-piece"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewKing(t *testing.T) {
	k := p.NewKing()
	assert.NotNil(t, k)
	assert.False(t, k.CanMoveMultiplePlaces())
	assert.Equal(t, 8, len(k.GetAllowedMoveDirections()))
}
