package game_engine

import (
	b "chess-board-simulator/board-layout"
	"sort"
)

type GameEngine struct {
	layout b.BoardLayout
}

func NewGameEngine() GameEngine {
	return GameEngine{layout: b.NewBoardLayout()}
}

func (ge GameEngine) Simulate(cell *b.Cell) []string {
	var possibleMoves []string
	piece := cell.GetPiece()
	for _, dir := range piece.GetAllowedMoveDirections() {
		newCell, canMove := ge.layout.MoveCell(cell, dir)

		for canMove {
			possibleMoves = append(possibleMoves, newCell.GetDisplayId())

			if !piece.CanMoveMultiplePlaces() {
				break
			}
			newCell, canMove = ge.layout.MoveCell(newCell, dir)
		}
	}
	sort.Strings(possibleMoves)
	return possibleMoves
}
