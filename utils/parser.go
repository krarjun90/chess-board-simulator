package utils

import (
	l "chess-board-simulator/board-layout"
	p "chess-board-simulator/board-piece"
	m "chess-board-simulator/messages"
	"errors"
	"strings"
)

func ParseInput(input string) (p.Piece, *l.Cell, error) {
	args := strings.Split(input, ",")
	if len(args) != 2 {
		return nil, nil, errors.New(m.InvalidInput)
	}
	piece, err := p.ParsePiece(strings.TrimSpace(args[0]))
	if err != nil {
		return nil, nil, err
	}

	cell, err := l.ParseSquare(strings.TrimSpace(args[1]))
	if err != nil {
		return nil, nil, err
	}
	return piece, cell, nil
}
