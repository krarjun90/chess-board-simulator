package board_layout

import (
	"errors"
	"fmt"
	"strconv"
)

type Cell struct {
	Row int
	Col int
}

func ParseSquare(cellIndex string) (*Cell, error) {
	if len(cellIndex) > 2 {
		return nil, errors.New(fmt.Sprintf("more than 2 characters: %v", cellIndex))
	}
	col := -1
	for i, v := range IndexArr {
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
	return &Cell{Row: int(Size - pos), Col: col}, nil
}
