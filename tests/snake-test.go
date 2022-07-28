package tests

import (
	"snake-game/components"
	"testing"
)

func newDoubleSnake(d components.Direction) *components.Snake {
	return components.NewSnake(d, []components.Cordinate{
		components.Cordinate{X: 1, Y: 0},
		components.Cordinate{X: 1, Y: 1},
		components.Cordinate{X: 1, Y: 2},
		components.Cordinate{X: 1, Y: 3},
		components.Cordinate{X: 1, Y: 4},
	})
}

func TestSnakeBodyMove(t *testing.T) {
	snake := newDoubleSnake(components.RIGHT)
	snake.Move()

	if snake.Body[0].X != 1 || snake.Body[0].Y != 1 {
		t.Fatalf("Invalid body position %x", snake.Body[0])
	}

	if snake.Body[1].X != 1 || snake.Body[1].Y != 2 {
		t.Fatalf("Invalid body position %x", snake.Body[1])
	}

	if snake.Body[2].X != 1 || snake.Body[2].Y != 3 {
		t.Fatalf("Invalid body position %x", snake.Body[2])
	}

	if snake.Body[3].X != 1 || snake.Body[3].Y != 4 {
		t.Fatalf("Invalid body position %x", snake.Body[3])
	}

	if snake.Body[4].X != 2 || snake.Body[4].Y != 4 {
		t.Fatalf("Invalid body position %x", snake.Body[4])
	}
}

func TestSnakeHeadMoveRight(t *testing.T) {
	snake := newDoubleSnake(components.RIGHT)
	snake.Move()

	if snake.Head().X != 2 || snake.Head().Y != 4 {
		t.Fatalf("Expected head to have moved to position [2 4], got %x", snake.Head())
	}
}

func TestSnakeHeadMoveUp(t *testing.T) {
	snake := newDoubleSnake(components.UP)
	snake.Move()

	if snake.Head().X != 1 || snake.Head().Y != 5 {
		t.Fatalf("Expected head to have moved to position [1 5], got %x", snake.Head())
	}
}

func TestSnakeHeadMoveDown(t *testing.T) {
	snake := newDoubleSnake(components.RIGHT)
	snake.Move()

	snake.ChangeDirection(components.DOWN)
	snake.Move()

	if snake.Head().X != 2 || snake.Head().Y != 3 {
		t.Fatalf("Expected head to have moved to position [2 3], got %x", snake.Head())
	}
}

func TestSnakeHeadMoveLeft(t *testing.T) {
	snake := newDoubleSnake(components.LEFT)
	snake.Move()

	if snake.Head().X != 0 || snake.Head().Y != 4 {
		t.Fatalf("Expected head to have moved to position [0 4], got %x", snake.Head())
	}
}

func TestChangeDirectionToNotOposity(t *testing.T) {
	snake := newDoubleSnake(components.DOWN)
	snake.ChangeDirection(components.RIGHT)
	if snake.Direction != components.RIGHT {
		t.Fatal("Expected to change Snake Direction to DOWN")
	}
}

func TestChangeDirectionToOposity(t *testing.T) {
	snake := newDoubleSnake(components.RIGHT)
	snake.ChangeDirection(components.LEFT)
	if snake.Direction == components.LEFT {
		t.Fatal("Expected not to have changed Snake Direction to LEFT")
	}
}

func TestChangeDirectionToInvalidDirection(t *testing.T) {
	snake := newDoubleSnake(components.RIGHT)
	snake.ChangeDirection(5)
	if snake.Direction != components.RIGHT {
		t.Fatal("Expected not to have changed Snake Direction")
	}
}

func TestSnakeDie(t *testing.T) {
	snake := newDoubleSnake(components.RIGHT)

	if err := snake.Die(); err.Error() != "Died" {
		t.Fatal("Expected Snake die() to return error")
	}
}

func TestSnakeDieWhenMoveOnTopOfItself(t *testing.T) {
	snake := newDoubleSnake(components.RIGHT)
	snake.Move()

	snake.ChangeDirection(components.DOWN)
	snake.Move()

	snake.ChangeDirection(components.LEFT)

	if err := snake.Die(); err != "Died" {
		t.Fatal("Expected Snake to die when moved on top of itself")
	}
}
