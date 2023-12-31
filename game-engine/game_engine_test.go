package game_engine_test

import (
	b "chess-board-simulator/board-layout"
	p "chess-board-simulator/board-piece"
	e "chess-board-simulator/game-engine"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

func TestSimulatePawnMovement(t *testing.T) {
	square, _ := b.ParseCell("G1")
	square.SetPiece(p.NewPawn())
	actualMoves := e.NewGameEngine().Simulate(square)
	assert.Equal(t, []string{"G2"}, actualMoves)
}

func TestSimulatePawnNoMovement(t *testing.T) {
	square, _ := b.ParseCell("G8")
	square.SetPiece(p.NewPawn())
	actualMoves := e.NewGameEngine().Simulate(square)
	assert.Equal(t, 0, len(actualMoves))
}

func TestSimulateKingMovement(t *testing.T) {
	// Taken From Test data
	exp := "C4, C5, C6, D4, D6, E4, E5, E6"
	expectedMoves := getSortedPlaces(exp)

	square, _ := b.ParseCell("D5")
	square.SetPiece(p.NewKing())
	actualMoves := e.NewGameEngine().Simulate(square)

	assert.Equal(t, expectedMoves, actualMoves)
}

func TestSimulateQueenMovement(t *testing.T) {
	// Taken From Test data
	exp := "A4, B4, C4, D4, F4, G4, H4, E1, E2, E3, E5, E6, E7, E8, A8, B7, C6, D5, F3, G2, H1, B1, C2, D3, F5, G6, H7"
	expectedMoves := getSortedPlaces(exp)

	square, _ := b.ParseCell("E4")
	square.SetPiece(p.NewQueen())
	actualMoves := e.NewGameEngine().Simulate(square)

	assert.Equal(t, expectedMoves, actualMoves)
}

func TestSimulateKnightMovement(t *testing.T) {
	exp := "C3, C5, D2, D6, F2, F6, G3, G5"
	expectedMoves := getSortedPlaces(exp)

	square, _ := b.ParseCell("E4")
	square.SetPiece(p.NewKnight())
	actualMoves := e.NewGameEngine().Simulate(square)

	sort.Strings(actualMoves)
	assert.Equal(t, expectedMoves, actualMoves)
}

func getSortedPlaces(exp string) []string {
	expArr := strings.Split(strings.ReplaceAll(exp, " ", ""), ",")
	sort.Strings(expArr)
	return expArr
}
