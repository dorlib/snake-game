package components

import "errors"

const (
	RIGHT direction = 1 + iota
	LEFT
	UP
	DOWN
)

type direction int

type snake struct {
	body      []cordinate
	direction direction
	length    int
}

func newSnake(d direction, b []cordinate) *snake {
	return &snake{
		body:      b,
		direction: d,
		length:    len(b),
	}
}

func (s *snake) changeDirection(d direction) {
	opposites := map[direction]direction{
		RIGHT: LEFT,
		LEFT:  RIGHT,
		UP:    DOWN,
		DOWN:  UP,
	}

	opp := opposites[d]

	if opp != 0 && opp != s.direction {
		s.direction = d
	}
}

func (s *snake) head() cordinate {
	return s.body[len(s.body)-1]
}

func (s *snake) die() error {
	return errors.New("game over")
}

func (s *snake) move() err {
	head := s.head()
	cord := cordinate{x: head.x, y: head.y}

	switch s.direction {
	case RIGHT:
		cord.x++
	case LEFT:
		cord.x--
	case UP:
		cord.y++
	case DOWN:
		cord.y--
	}

	if s.isOnPosition(cord) {
		return s.die()
	}

	if s.length > len(s.body) {
		s.body = append(s.body, cord)
	} else {
		s.body = append(s.body[1:], cord)
	}
	return nil
}

func (s *snake) isOnPosition(c cordinate) bool {
	for _, b := range s.body {
		if b.x == c.x && b.y == c.y {
			return true
		}
	}
	return false
}
