package components

import (
	"math/rand"
	"time"
)

type arena struct {
	food       *Food
	snake      *snake
	hasFood    func(*arena, coord) bool
	height     int
	width      int
	pointsChan chan int
}

func newArena(s *snake, p chan int, h int, w int) *arena {
	rand.Seed(time.Now().UnixNano())

	a := &arena{
		snake:      s,
		height:     h,
		width:      w,
		pointsChan: p,
		hasFood:    hasFood,
	}

	a.placeFood()
	return a
}

func (a *arena) placeFood() {
	var x, y int

	for {
		x = rand.Intn(a.width)
		y = rand.Intn(a.height)

		if !a.isOccupied(cordinate{x: x, y: y}) {
			break
		}
	}
}

func (a *arena) isOccupied(cord cordinate) bool {
	return a.snake.isOnPosition(cord)
}

func (a *arena) snakeOutOfArea() bool {
	h := a.snake.head()
	return h.x > a.width || h.y > a.height || h.x < 0 || h.y < 0
}

func (a *arena) addScore(p int) {
	a.pointsChan <- p
}

func (a *arena) moveSnake() err {
	if err := a.snake.move(); err != nil {
		return err
	}

	if a.snakeOutOfArea() {
		return a.snake.die()
	}

	if a.hasFood(a, a.snake.head()) {
		go a.addScore(a.food.Score)
		a.snake.length++
		a.placeFood()
	}
	return nil
}

func hasFood(a *arena, cord cordinate) bool {
	return cord.x == a.food.x && cord.y == a.food.y
}
