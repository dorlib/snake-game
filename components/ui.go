package components

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
	snakeColor   = termbox.ColorGreen
)

func renderFood(left int, bottom int, f *Food) {
	termbox.SetCell(left+f.X, bottom-f.Y, f.Emoji, defaultColor, bgColor)
}

func renderSnake(left, bottom int, s *Snake) {
	for _, b := range s.Body {
		termbox.SetCell(left+b.x, bottom-b.y, ' ', snakeColor, snakeColor)
	}
}

func renderArena(a *Arena, top, bottom, left int) {
	for i := top; i < bottom; i++ {
		termbox.SetCell(left-1, i, '|', defaultColor, bgColor)
		termbox.SetCell(left+a.width, i, '|', defaultColor, bgColor)
	}

	termbox.SetCell(left-1, top, '┌', defaultColor, bgColor)
	termbox.SetCell(left-1, bottom, '└', defaultColor, bgColor)
	termbox.SetCell(left+a.width, top, '┐', defaultColor, bgColor)
	termbox.SetCell(left+a.width, bottom, '┘', defaultColor, bgColor)

	fill(left, top, a.width, 1, termbox.Cell{Ch: '─'})
	fill(left, bottom, a.width, 1, termbox.Cell{Ch: '─'})
}

func renderTitle(left, top int) {
	tbprint(left, top-1, defaultColor, defaultColor, "Snake Game")
}

func renderQuitMessage(right int, bottom int) {
	m := "Press ESC to quit"
	tbprint(right-17, bottom+1, defaultColor, defaultColor, m)
}

func renderScore(left, bottom, s int) {
	score := fmt.Sprintf("Score: %v", s)
	tbprint(left, bottom+1, defaultColor, defaultColor, score)
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < h; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}

}

func tbprint(x int, y int, fg termbox.Attribute, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func (g *game) Render() error {
	termbox.Clear(defaultColor, defaultColor)

	var (
		w, h   = termbox.Size()
		midY   = h / 2
		left   = (w - g.Arena.width) / 2
		right  = (h - g.Arena.height) / 2
		top    = midY - (g.Arena.height / 2)
		bottom = midY - (g.Arena.height / 2) + 1
	)

	renderTitle(left, top)
	renderArena(g.Arena, top, bottom, left)
	renderSnake(left, bottom, g.Arena.Snake)
	renderFood(left, bottom, g.Arena.Food)
	renderScore(left, bottom, g.Score)
	renderQuitMessage(right, bottom)

	return termbox.Flush()
}
