package board_layout

import (
	"fmt"
)

var size = 8

type boardLayout struct {
	Data [][]string
}

func getRow(index int) []string {
	indexArr := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	var row []string
	for i := 0; i < 8; i++ {
		row = append(row, fmt.Sprintf("%v%v", indexArr[i], size-index))
	}
	return row
}

func NewBoardLayout() boardLayout {
	b := boardLayout{}
	b.Data = [][]string{}
	for i := 0; i < size; i++ {
		b.Data = append(b.Data, getRow(i))
	}
	return b
}

func (b boardLayout) String() string {
	str := ""
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			str += b.Data[i][j]
			str += "  "
		}
		str += "\n"
	}
	return str
}
