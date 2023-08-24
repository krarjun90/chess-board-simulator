package utils

import (
	"errors"
	"strings"
	l "technogise/chess-board/board-layout"
	p "technogise/chess-board/board-piece"
	"technogise/chess-board/messages"
)

func ParseInput(input string) (p.Piece, *l.Cell, error) {
	args := strings.Split(input, ",")
	if len(args) != 2 {
		return nil, nil, errors.New(messages.InvalidInput)
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
