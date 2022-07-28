package components

import "errors"

const (
	RIGHT Direction = 1 + iota
	LEFT
	UP
	DOWN
)

type Direction int

type Snake struct {
	Body      []Cordinate
	Direction Direction
	Length    int
}

func NewSnake(d Direction, b []Cordinate) *Snake {
	return &Snake{
		Body:      b,
		Direction: d,
		Length:    len(b),
	}
}

func (s *Snake) ChangeDirection(d Direction) {
	opposites := map[Direction]Direction{
		RIGHT: LEFT,
		LEFT:  RIGHT,
		UP:    DOWN,
		DOWN:  UP,
	}

	opp := opposites[d]

	if opp != 0 && opp != s.Direction {
		s.Direction = d
	}
}

func (s *Snake) Head() Cordinate {
	return s.Body[len(s.Body)-1]
}

func (s *Snake) Die() error {
	return errors.New("game over")
}

func (s *Snake) Move() err {
	head := s.Head()
	cord := Cordinate{X: head.X, Y: head.Y}

	switch s.Direction {
	case RIGHT:
		cord.X++
	case LEFT:
		cord.X--
	case UP:
		cord.Y++
	case DOWN:
		cord.Y--
	}

	if s.isOnPosition(cord) {
		return s.Die()
	}

	if s.Length > len(s.Body) {
		s.Body = append(s.Body, cord)
	} else {
		s.Body = append(s.Body[1:], cord)
	}
	return nil
}

func (s *Snake) isOnPosition(c Cordinate) bool {
	for _, b := range s.Body {
		if b.X == c.X && b.Y == c.Y {
			return true
		}
	}
	return false
}
