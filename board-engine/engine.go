package board_engine

import (
	"sort"
	b "technogise/chess-board/board-layout"
	p "technogise/chess-board/board-piece"
)

func SimulateMovement(layout b.BoardLayout, piece p.Piece, cell b.Cell) []string {
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
			possibleMoves = append(possibleMoves, layout.Data[row][col])

			if !piece.CanMoveMultiplePlaces() {
				break
			}
		}
	}
	sort.Strings(possibleMoves)
	return possibleMoves
}
