package components

import (
	"math/rand"
	"time"
)

type Arena struct {
	Food      *Food
	Snake     *Snake
	HasFood   func(*Arena, Cordinate) bool
	height    int
	width     int
	scoreChan chan int
}

func NewArena(s *Snake, p chan int, h int, w int) *Arena {
	rand.Seed(time.Now().UnixNano())

	a := &Arena{
		Snake:     s,
		height:    h,
		width:     w,
		scoreChan: p,
		HasFood:   HasFood,
	}

	a.placeFood()
	return a
}

func (a *Arena) placeFood() {
	var x, y int

	for {
		x = rand.Intn(a.width)
		y = rand.Intn(a.height)

		if !a.isOccupied(Cordinate{X: x, Y: y}) {
			break
		}
	}
}

func (a *Arena) isOccupied(cord Cordinate) bool {
	return a.Snake.isOnPosition(cord)
}

func (a *Arena) snakeOutOfArea() bool {
	h := a.Snake.Head()
	return h.X > a.width || h.Y > a.height || h.X < 0 || h.Y < 0
}

func (a *Arena) addScore(p int) {
	a.scoreChan <- p
}

func (a *Arena) MoveSnake() err {
	if err := a.Snake.Move(); err != nil {
		return err
	}

	if a.snakeOutOfArea() {
		return a.Snake.Die()
	}

	if a.HasFood(a, a.Snake.Head()) {
		go a.addScore(a.Food.Score)
		a.Snake.Length++
		a.placeFood()
	}
	return nil
}

func HasFood(a *Arena, cord Cordinate) bool {
	return cord.X == a.Food.X && cord.Y == a.Food.Y
}
