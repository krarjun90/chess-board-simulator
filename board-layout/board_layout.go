package board_layout

import (
	"fmt"
)

const Size = 8

var IndexArr = []string{"A", "B", "C", "D", "E", "F", "G", "H"}

type BoardLayout struct {
	data [][]string
}

func (b BoardLayout) GetLayout() [][]string {
	return b.data
}

func getRow(index int) []string {
	var row []string
	for i := 0; i < 8; i++ {
		row = append(row, fmt.Sprintf("%v%v", IndexArr[i], Size-index))
	}
	return row
}

func NewBoardLayout() BoardLayout {
	b := BoardLayout{}
	b.data = [][]string{}
	for i := 0; i < Size; i++ {
		b.data = append(b.data, getRow(i))
	}
	return b
}

func (b BoardLayout) String() string {
	str := ""
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			str += b.data[i][j]
			str += "  "
		}
		str += "\n"
	}
	return str
}
