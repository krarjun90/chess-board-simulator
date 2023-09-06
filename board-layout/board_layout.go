package board_layout

import (
	board_piece "chess-board-simulator/board-piece"
	"chess-board-simulator/consts"
)

type BoardLayout struct {
	data map[int][]*Cell
}

func (b BoardLayout) GetLayout() map[int][]*Cell {
	return b.data
}

func getRow(rankIndex int) []*Cell {
	var row []*Cell
	for fileIndex := 0; fileIndex < consts.GridSize; fileIndex++ {
		row = append(row, NewCell(fileIndex, rankIndex))
	}
	return row
}

func NewBoardLayout() BoardLayout {
	b := BoardLayout{}
	b.data = make(map[int][]*Cell, consts.GridSize)
	for i := 0; i < consts.GridSize; i++ {
		b.data[i] = getRow(i)
	}
	return b
}

func (b BoardLayout) MoveCell(cell *Cell, direction board_piece.Move) (*Cell, bool) {
	newFile := cell.GetFileIndex() + direction.GetHPos()
	newRank := cell.GetRankIndex() - direction.GetVPos()
	if cellWithInGrid(newRank, newFile) {
		return b.data[newRank][newFile], true
	}
	return nil, false
}

func cellWithInGrid(row, col int) bool {
	return col >= 0 && col < consts.GridSize && row >= 0 && row < consts.GridSize
}

func (b BoardLayout) String() string {
	str := ""
	for i := 0; i < consts.GridSize; i++ {
		for j := 0; j < consts.GridSize; j++ {
			str += b.data[i][j].GetDisplayId()
			str += "  "
		}
		str += "\n"
	}
	return str
}
