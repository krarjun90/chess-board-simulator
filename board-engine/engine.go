package board_engine

import (
	b "chess-board-simulator/board-layout"
	p "chess-board-simulator/board-piece"
	"sort"
)

type Engine struct {
	layout b.BoardLayout
}

func NewEngine() Engine {
	return Engine{layout: b.NewBoardLayout()}
}

func (s Engine) Simulate(piece p.Piece, cell b.Cell) []string {
	var possibleMoves []string
	for _, dir := range piece.GetAllowedMoveDirections() {
		row := cell.Row
		col := cell.Col

		for true {
			row = row - dir.GetVPos()
			col = col - dir.GetHPos()

			if (col < 0 || col > 7) || (row < 0 || row > 7) {
				break
			}
			possibleMoves = append(possibleMoves, s.layout.GetLayout()[row][col])

			if !piece.CanMoveMultiplePlaces() {
				break
			}
		}
	}
	sort.Strings(possibleMoves)
	return possibleMoves
}
