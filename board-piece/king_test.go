package board_piece_test

import (
	"github.com/stretchr/testify/assert"
	p "technogise/chess-board/board-piece"
	"testing"
)

func TestNewKing(t *testing.T) {
	k := p.NewKing()
	assert.NotNil(t, k)
	assert.False(t, k.CanMoveMultiplePlaces())
	assert.Equal(t, 8, len(k.GetAllowedMoveDirections()))
}
