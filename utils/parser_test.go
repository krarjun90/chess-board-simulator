package utils_test

import (
	"github.com/stretchr/testify/assert"
	l "technogise/chess-board/board-layout"
	p "technogise/chess-board/board-piece"
	"technogise/chess-board/utils"
	"testing"
)

func TestParseInputInvalidFormat(t *testing.T) {
	piece, cell, err := utils.ParseInput("King, A1, Queen")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid format", err.Error())
	assert.Nil(t, piece)
	assert.Nil(t, cell)
}

func TestParseInputInvalidPiece(t *testing.T) {
	piece, cell, err := utils.ParseInput("Kings, A1")
	assert.NotNil(t, err)
	assert.Equal(t, "piece not supported", err.Error())
	assert.Nil(t, piece)
	assert.Nil(t, cell)
}

func TestParseInputInvalidCell(t *testing.T) {
	piece, cell, err := utils.ParseInput("King, A0")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid row", err.Error())
	assert.Nil(t, piece)
	assert.Nil(t, cell)
}

func TestParseInputSuccess(t *testing.T) {
	piece, cell, err := utils.ParseInput("King, A1")
	assert.Nil(t, err)
	assert.Equal(t, p.NewKing(), piece)
	assert.Equal(t, l.Cell{Row: 7}, *cell)
}
