package tests

import (
	"snake-game/components"
	_ "snake-game/components"
	"testing"
)

var pointsDouble = make(chan int)

func newDoubleArenaWithFoodFinder(h, w int, f func(*components.Arena, components.Cordinate) bool) *components.Arena {
	a := newDoubleArena(h, w)
	a.HasFood = f
	return a
}

func newDoubleArena(h, w int) *components.Arena {
	s := components.NewSnake(components.RIGHT, []components.Cordinate{
		components.Cordinate{X: 1, Y: 0},
		components.Cordinate{X: 1, Y: 1},
		components.Cordinate{X: 1, Y: 2},
		components.Cordinate{X: 1, Y: 3},
		components.Cordinate{X: 1, Y: 4},
	})

	return components.NewArena(s, pointsDouble, h, w)
}

func TestArenaHaveFoodPlaced(t *testing.T) {
	if a := newDoubleArena(20, 20); a.Food == nil {
		t.Fatal("Arena expected to have food placed")
	}
}

func TestMoveSnakeOutOfArenaHeightLimit(t *testing.T) {
	a := newDoubleArena(4, 10)
	a.Snake.ChangeDirection(UP)

	if err := a.MoveSnake(); err == nil || err.Error() != "Died" {
		t.Fatal("Expected Snake to die when moving outside the Arena height limits")
	}
}

func TestMoveSnakeOutOfArenaWidthLimit(t *testing.T) {
	a := newDoubleArena(10, 1)
	a.Snake.ChangeDirection(components.LEFT)

	if err := a.MoveSnake(); err == nil || err.Error() != "Died" {
		t.Fatal("Expected Snake to die when moving outside the Arena height limits")
	}
}

func TestPlaceNewFoodWhenEatFood(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*components.Arena, coord) bool {
		return true
	})

	f := a.Food

	a.MoveSnake()

	if a.Food.X == f.X && a.Food.Y == f.Y {
		t.Fatal("Expected new food to have been placed on Arena")
	}
}

func TestIncreaseSnakeLengthWhenEatFood(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*components.Arena, components.Cordinate) bool {
		return true
	})

	l := a.Snake.Length

	a.MoveSnake()

	if a.Snake.Length != l+1 {
		t.Fatal("Expected Snake to have grown")
	}
}

func TestAddPointsWhenEatFood(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*components.Arena, components.Cordinate) bool {
		return true
	})

	if p, ok := <-pointsDouble; ok && p != a.Food.Score {
		t.Fatalf("Value %d was expected but got %d", a.Food.Score, p)
	}

	a.MoveSnake()
}

func TestDoesNotAddPointsWhenFoodNotFound(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*components.Arena, components.Cordinate) bool {
		return false
	})

	select {
	case p, _ := <-components.ScoreChan:
		t.Fatalf("No point was expected to be received but received %d", p)
	default:
		close(components.ScoreChan)
	}

	a.MoveSnake()
}

func TestDoesNotPlaceNewFoodWhenFoodNotFound(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*components.Arena, components.Cordinate) bool {
		return false
	})

	f := a.Food

	a.MoveSnake()

	if a.Food.X != f.X || a.Food.Y != f.Y {
		t.Fatal("Food in Arena expected not to have changed")
	}
}

func TestDoesNotIncreaseSnakeLengthWhenFoodNotFound(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*components.Arena, components.Cordinate) bool {
		return false
	})

	l := a.Snake.Length

	a.MoveSnake()

	if a.Snake.Length != l {
		t.Fatal("Expected Snake not to have grown")
	}
}

func TestHasFood(t *testing.T) {
	a := newDoubleArena(20, 20)

	if !components.HasFood(a, components.Cordinate{X: a.Food.X, Y: a.Food.Y}) {
		t.Fatal("Food expected to be found")
	}
}

func TestHasNotFood(t *testing.T) {
	a := newDoubleArena(20, 20)

	if components.HasFood(a, components.Cordinate{X: a.Food.X - 1, Y: a.Food.Y}) {
		t.Fatal("No food expected to be found")
	}
}
