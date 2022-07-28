package tests

import (
	"snake-game/components"
	"testing"
	"time"
)

func TestDefaultGameScore(t *testing.T) {
	g := components.NewGame()

	if g.Score != 0 {
		t.Fatalf("Initial Game Score expected to be 0 but it was %d", g.Score)
	}
}

func TestGameMoveInterval(t *testing.T) {
	e := time.Duration(85) * time.Millisecond
	g := components.NewGame()
	g.Score = 150

	if d := g.MoveInterval(); d != e {
		t.Fatalf("Expected move interval to be %d but got %d", e, d)
	}
}

func TestAddPoints(t *testing.T) {
	g := components.NewGame()
	s := g.Score
	g.AddScore(10)

	if s != 0 || g.Score != 10 {
		t.Fatal("Expected 10 points to have been added to Game Score")
	}
}

func TestRetryGoBackToGameInitialState(t *testing.T) {
	g := components.NewGame()
	initScore := g.Score
	initSnake := g.Arena.Snake

	g.Arena.Snake.ChangeDirection(components.UP)
	g.Arena.MoveSnake()
	g.AddScore(10)
	g.End()

	g.Retry()

	if g.Score != initScore {
		t.Fatal("Expected Score to have been reset")
	}

	for i, c := range g.Arena.Snake.Body {
		if initSnake.Body[i].X == c.X && initSnake.Body[i].Y == c.Y {
			t.Fatal("Expected Snake body to have been reset")
		}
	}

	if g.Arena.Snake.Direction == initSnake.Direction {
		t.Fatal("Expected Snake direction to have been reset")
	}
}
