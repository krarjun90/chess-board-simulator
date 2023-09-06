package board_layout

import (
	p "chess-board-simulator/board-piece"
	"chess-board-simulator/consts"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Cell struct {
	fileIndex int
	rankIndex int
	piece     p.Piece
}

func (cell *Cell) GetFileIndex() int {
	return cell.fileIndex
}

func (cell *Cell) GetRankIndex() int {
	return cell.rankIndex
}

func (cell *Cell) GetPiece() p.Piece {
	return cell.piece
}

func (cell *Cell) SetPiece(piece p.Piece) {
	cell.piece = piece
}

func (cell *Cell) GetDisplayId() string {
	return fmt.Sprintf("%v%v", consts.FileIds[cell.GetFileIndex()], consts.GridSize-cell.GetRankIndex())
}

func NewCell(fileIndex, rankIndex int) *Cell {
	return &Cell{fileIndex: fileIndex, rankIndex: rankIndex}
}

func ParseCell(cellIndex string) (*Cell, error) {
	if len(cellIndex) > 2 {
		return nil, fmt.Errorf("more than 2 characters: %v", cellIndex)
	}
	fileIndex, err := getFileIndex(cellIndex)
	if err != nil {
		return nil, err
	}

	rankIndex, err := getRankIndex(cellIndex, err)
	if err != nil {
		return nil, err
	}
	return &Cell{fileIndex: fileIndex, rankIndex: rankIndex}, nil
}

func getRankIndex(cellIndex string, err error) (int, error) {
	pos, err := strconv.ParseInt(string(cellIndex[1]), 10, 64)
	if err != nil {
		return int(pos), errors.New("invalid row")
	}
	if pos <= 0 || pos > 8 {
		return int(pos), errors.New("invalid row")
	}
	// A1-H1 is in row 7
	return int(consts.GridSize - pos), nil
}

func getFileIndex(cellIndex string) (int, error) {
	fileIndex := -1
	for i, v := range consts.FileIds {
		if v == strings.ToUpper(string(cellIndex[0])) {
			fileIndex = i
			break
		}
	}
	if fileIndex == -1 {
		return fileIndex, errors.New("invalid column")
	}
	return fileIndex, nil
}
