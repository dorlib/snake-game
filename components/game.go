package components

import (
	"github.com/nsf/termbox-go"
	"time"
)

var (
	ScoreChan          = make(chan int)
	keyboardEventsChan = make(chan keyboardEvent)
)

type game struct {
	Arena  *Arena
	Score  int
	isOver bool
}

func initialSnake() *Snake {
	return NewSnake(RIGHT, []Cordinate{
		Cordinate{X: 1, Y: 1},
		Cordinate{X: 1, Y: 2},
		Cordinate{X: 1, Y: 3},
		Cordinate{X: 1, Y: 4},
	})
}

func initialScore() int {
	return 0
}

func initialArena() *Arena {
	return newArena(initialSnake(), ScoreChan, 20, 50)
}

func (g *game) End() {
	g.isOver = true
}

func (g *game) MoveInterval() time.Duration {
	ms := 100 - (g.Score / 10)
	return time.Duration(ms) * time.Millisecond
}
func (g *game) Retry() {
	g.Arena = initialArena()
	g.Score = initialScore()
	g.isOver = false
}

func (g *game) AddScore(p int) {
	g.Score += p
}

// creating new game object
func NewGame() *game {
	return &game{Arena: initialArena(), Score: initialScore(), isOver: false}
}

// start the new created game object
func (g *game) start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	go listenToKeyboard(keyboardEventsChan)

	if err := g.Render(); err != nil {
		panic(err)
	}

mainloop:
	for {
		select {
		case p := <-ScoreChan:
			g.AddScore(p)
		case e := <-keyboardEventsChan:
			switch e.eventType {
			case MOVE:
				d := KeyToDirection(e.key)
				g.Arena.Snake.ChangeDirection(d)
			case RETRY:
				g.Retry()
			case END:
				break mainloop
			}
		default:
			if !g.isOver {
				if err := g.Arena.MoveSnake(); err != nil {
					g.End()
				}
			}
			if err := g.Render(); err != nil {
				panic(err)
			}
			time.Sleep(g.MoveInterval())
		}
	}
}
