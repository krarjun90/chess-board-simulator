package board_piece_test

import (
	"github.com/stretchr/testify/assert"
	p "technogise/chess-board/board-piece"
	"testing"
)

func TestNewQueen(t *testing.T) {
	q := p.NewQueen()
	assert.NotNil(t, q)
	assert.True(t, q.CanMoveMultiplePlaces())
	assert.Equal(t, 8, len(q.GetAllowedMoveDirections()))
}
