package components

import (
	"github.com/nsf/termbox-go"
	"time"
)

var (
	scoreChan          = make(chan int)
	keyboardEventsChan = make(chan keyboardEvent)
)

type game struct {
	arena  *arena
	score  int
	isOver bool
}

func initialSnake() *snake {
	return newSnake(RIGHT, []cordinate{
		cordinate{x: 1, y: 1},
		cordinate{x: 1, y: 2},
		cordinate{x: 1, y: 3},
		cordinate{x: 1, y: 4},
	})
}

func initialScore() int {
	return 0
}

func initialArena() *arena {
	return newArena(initialSnake(), scoreChan, 20, 50)
}

func (g *game) end() {
	g.isOver = true
}

func (g *game) moveInterval() time.Duration {
	ms := 100 - (g.score / 10)
	return time.Duration(ms) * time.Millisecond
}
func (g *game) retry() {
	g.arena = initialArena()
	g.score = initialScore()
	g.isOver = false
}

func (g *game) addScore(p int) {
	g.score += p
}

// creating new game object
func newGame() * {
	return &game{arena: initialArena(), score: initialScore(), isOver: false}
}

// start the new created game object
func (g *game) start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	go listenToKeyboard(keyboardEventsChan)

	if err := g.render(); err != nil {
		panic(err)
	}

	mainloop:
		for {
			select{
			case p := <- scoreChan:
				g.addScore(p)
			case e := <- keyboardEventsChan:
				switch e.eventType {
				case MOVE:
					d := KeyToDirection(e.key)
					g.arena.snake.changeDirection(d)
				case RETRY:
					g.retry()
				case END:
					break mainloop
				}
			default:
				if !g.isOver {
					if err := g.arena.moveSnake(); err != nil {
						g.end()
					}
				}
				if err := g.render(); err != nil {
					panic(err)
				}
				time.Sleep(g.moveInterval())
			}
		}
}