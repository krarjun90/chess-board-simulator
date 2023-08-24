package board_engine_test

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	engine "technogise/chess-board/board-engine"
	l "technogise/chess-board/board-layout"
	p "technogise/chess-board/board-piece"
	"testing"
)

func TestSimulatePawnMovement(t *testing.T) {
	layout := l.NewBoardLayout()
	square, _ := l.ParseSquare("G1")
	movement := engine.SimulateMovement(layout, p.NewPawn(), *square)
	assert.Equal(t, []string{"G2"}, movement)
}

func TestSimulateKingMovement(t *testing.T) {
	// Taken From Test Data
	exp := "C4, C5, C6, D4, D6, E4, E5, E6"
	expectedMoves := getSortedPlaces(exp)

	layout := l.NewBoardLayout()
	square, _ := l.ParseSquare("D5")
	actualMoves := engine.SimulateMovement(layout, p.NewKing(), *square)

	sort.Strings(actualMoves)
	assert.Equal(t, expectedMoves, actualMoves)
}

func TestSimulateQueenMovement(t *testing.T) {
	// Taken From Test Data
	exp := "A4, B4, C4, D4, F4, G4, H4, E1, E2, E3, E5, E6, E7, E8, A8, B7, C6, D5, F3, G2, H1, B1, C2, D3, F5, G6, H7"
	expectedMoves := getSortedPlaces(exp)

	layout := l.NewBoardLayout()
	square, _ := l.ParseSquare("E4")
	actualMoves := engine.SimulateMovement(layout, p.NewQueen(), *square)

	sort.Strings(actualMoves)
	assert.Equal(t, expectedMoves, actualMoves)
}

func getSortedPlaces(exp string) []string {
	expArr := strings.Split(strings.ReplaceAll(exp, " ", ""), ",")
	sort.Strings(expArr)
	return expArr
}