package board_engine_test

import (
	"github.com/stretchr/testify/assert"
	engine "technogise/chess-board/board-engine"
	l "technogise/chess-board/board-layout"
	p "technogise/chess-board/board-piece"
	"testing"
)

func TestSimulateMovement(t *testing.T) {
	layout := l.NewBoardLayout()
	movement := engine.SimulateMovement(layout, p.NewPawn(), l.Cell{Row: 1, Col: 1})
	assert.Equal(t, []string{layout.Data[0][1]}, movement)
}
