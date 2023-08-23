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
