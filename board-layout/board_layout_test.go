package board_layout_test

import (
	"github.com/stretchr/testify/assert"
	layout "technogise/chess-board/board-layout"
	"testing"
)

func TestNewBoardLayout(t *testing.T) {
	b := layout.NewBoardLayout()
	assert.NotNil(t, b)
	assert.NotNil(t, b.Data)
	assert.Equal(t, 8, len(b.Data))
	for i := 0; i < 8; i++ {
		assert.Equal(t, 8, len(b.Data[i]))
	}
}

func TestValidateCell(t *testing.T) {
	cell, err := layout.ParseSquare("A1")
	assert.Nil(t, err)
	assert.NotNil(t, cell)
	assert.Equal(t, 7, cell.Row)
	assert.Equal(t, 0, cell.Col)

	cell, err = layout.ParseSquare("A8")
	assert.Nil(t, err)
	assert.NotNil(t, cell)
	assert.Equal(t, 0, cell.Row)
	assert.Equal(t, 0, cell.Col)

	cell, err = layout.ParseSquare("H5")
	assert.Nil(t, err)
	assert.NotNil(t, cell)
	assert.Equal(t, 3, cell.Row)
	assert.Equal(t, 7, cell.Col)
}

func TestValidateInvalidRow(t *testing.T) {
	cell, err := layout.ParseSquare("H0")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid row", err.Error())
	assert.Nil(t, cell)

	cell, err = layout.ParseSquare("H9")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid row", err.Error())
	assert.Nil(t, cell)

	cell, err = layout.ParseSquare("HH")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid row", err.Error())
	assert.Nil(t, cell)
}

func TestValidateInvalidColumn(t *testing.T) {
	cell, err := layout.ParseSquare("X0")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid column", err.Error())
	assert.Nil(t, cell)

	cell, err = layout.ParseSquare("11")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid column", err.Error())
	assert.Nil(t, cell)
}

func TestValidatePositionFormat(t *testing.T) {
	cell, err := layout.ParseSquare("X11")
	assert.NotNil(t, err)
	assert.Equal(t, "more than 2 characters", err.Error())
	assert.Nil(t, cell)
}
