package board_layout

import (
	"errors"
	"fmt"
	"strconv"
)

const size = 8

var indexArr = []string{"A", "B", "C", "D", "E", "F", "G", "H"}

type BoardLayout struct {
	Data [][]string
}

type Cell struct {
	Row int
	Col int
}

func getRow(index int) []string {
	var row []string
	for i := 0; i < 8; i++ {
		row = append(row, fmt.Sprintf("%v%v", indexArr[i], size-index))
	}
	return row
}

func NewBoardLayout() BoardLayout {
	b := BoardLayout{}
	b.Data = [][]string{}
	for i := 0; i < size; i++ {
		b.Data = append(b.Data, getRow(i))
	}
	return b
}

func (b BoardLayout) String() string {
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

func ParseSquare(cellIndex string) (*Cell, error) {
	if len(cellIndex) > 2 {
		return nil, errors.New("more than 2 characters")
	}
	col := -1
	for i, v := range indexArr {
		if v == string(cellIndex[0]) {
			col = i
			break
		}
	}
	if col == -1 {
		return nil, errors.New("invalid column")
	}

	pos, err := strconv.ParseInt(string(cellIndex[1]), 10, 64)
	if err != nil {
		return nil, errors.New("invalid row")
	}

	if pos <= 0 || pos > 8 {
		return nil, errors.New("invalid row")
	}
	return &Cell{Row: int(size - pos), Col: col}, nil
}
