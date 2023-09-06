package board_layout_test

import (
	b "chess-board-simulator/board-layout"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateCell(t *testing.T) {
	cell, err := b.ParseCell("A1")
	assert.Nil(t, err)
	assert.NotNil(t, cell)
	assert.Equal(t, 7, cell.GetRankIndex())
	assert.Equal(t, 0, cell.GetFileIndex())

	cell, err = b.ParseCell("A8")
	assert.Nil(t, err)
	assert.NotNil(t, cell)
	assert.Equal(t, 0, cell.GetRankIndex())
	assert.Equal(t, 0, cell.GetFileIndex())

	cell, err = b.ParseCell("H5")
	assert.Nil(t, err)
	assert.NotNil(t, cell)
	assert.Equal(t, 3, cell.GetRankIndex())
	assert.Equal(t, 7, cell.GetFileIndex())
}

func TestValidateInvalidRow(t *testing.T) {
	cell, err := b.ParseCell("H0")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid row", err.Error())
	assert.Nil(t, cell)

	cell, err = b.ParseCell("H9")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid row", err.Error())
	assert.Nil(t, cell)

	cell, err = b.ParseCell("HH")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid row", err.Error())
	assert.Nil(t, cell)
}

func TestValidateInvalidColumn(t *testing.T) {
	cell, err := b.ParseCell("X0")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid column", err.Error())
	assert.Nil(t, cell)

	cell, err = b.ParseCell("11")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid column", err.Error())
	assert.Nil(t, cell)
}

func TestValidatePositionFormat(t *testing.T) {
	cell, err := b.ParseCell("X11")
	assert.NotNil(t, err)
	assert.Equal(t, "more than 2 characters: X11", err.Error())
	assert.Nil(t, cell)
}
