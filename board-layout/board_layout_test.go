package board_layout_test

import (
	layout "chess-board-simulator/board-layout"
	board_piece "chess-board-simulator/board-piece"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBoardLayout(t *testing.T) {
	b := layout.NewBoardLayout()
	fmt.Println(b)

	assert.NotNil(t, b)
	assert.NotNil(t, b.GetLayout())
	assert.Equal(t, 8, len(b.GetLayout()))
	for i := 0; i < 8; i++ {
		assert.Equal(t, 8, len(b.GetLayout()[i]))
	}
}

func TestMoveCell_CanNotMove(t *testing.T) {
	l := layout.NewBoardLayout()
	newCell, canMove := l.MoveCell(layout.NewCell(7, 7), board_piece.NewMove(-1, -1))
	assert.False(t, canMove)
	assert.Nil(t, newCell)
}

func TestMoveCell_CanMove(t *testing.T) {
	l := layout.NewBoardLayout()
	newCell, canMove := l.MoveCell(layout.NewCell(7, 7), board_piece.NewMove(-1, 1))
	assert.True(t, canMove)
	assert.NotNil(t, newCell)
}
