package board_engine

import boardLayout "technogise/chess-board/board-layout"
import p "technogise/chess-board/board-piece"

func SimulateMovement(layout boardLayout.BoardLayout, piece p.Piece, cell boardLayout.Cell) []string {
	var possibleMoves []string
	switch piece.(type) {
	case p.Pawn:
		possibleMoves = simulatePawnMovement(layout, cell)
	default:
		return []string{}
	}
	return possibleMoves
}

func simulatePawnMovement(layout boardLayout.BoardLayout, cell boardLayout.Cell) []string {
	if cell.Row == 0 {
		return []string{}
	}
	newRow := cell.Row - 1
	return []string{layout.Data[newRow][cell.Col]}
}
