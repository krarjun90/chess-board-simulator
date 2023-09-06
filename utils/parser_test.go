package utils_test

import (
	p "chess-board-simulator/board-piece"
	"chess-board-simulator/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInputInvalidFormat(t *testing.T) {
	cell, err := utils.ParseInput("King, A1, Queen")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid format", err.Error())
	assert.Nil(t, cell)
}

func TestParseInputInvalidPiece(t *testing.T) {
	cell, err := utils.ParseInput("Kings, A1")
	assert.NotNil(t, err)
	assert.Equal(t, "piece not supported", err.Error())
	assert.Nil(t, cell)
}

func TestParseInputInvalidCell(t *testing.T) {
	cell, err := utils.ParseInput("King, A0")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid row", err.Error())
	assert.Nil(t, cell)
}

func TestParseInputSuccess(t *testing.T) {
	cell, err := utils.ParseInput("King, A1")
	assert.Nil(t, err)
	assert.Equal(t, p.NewKing(), cell.GetPiece())
	assert.Equal(t, 7, cell.GetRankIndex())
	assert.Equal(t, 0, cell.GetFileIndex())
}
