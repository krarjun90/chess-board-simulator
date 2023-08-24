package board_engine

import (
	boardLayout "technogise/chess-board/board-layout"
)
import p "technogise/chess-board/board-piece"

func SimulateMovement(layout boardLayout.BoardLayout, piece p.Piece, cell boardLayout.Cell) []string {
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
	return possibleMoves
}
