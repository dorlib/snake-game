package components

import (
	"math/rand"
	"os"
	"strings"
)

type food struct {
	emoji rune
	x     int
	y     int
	score int
}

func newFood(x int, y int) *food {
	return &food{
		emoji: getFoodEmoji(),
		x:     x,
		y:     y,
		score: 10,
	}
}

func getFoodEmoji() rune {
	if hasUnicodeSupport() {
		return randomEmojies()
	} else {
		return '@'
	}
}

func randomEmojies() rune {
	emojies := []rune{
		'🍒',
		'🍍',
		'🍑',
		'🍇',
		'🍏',
		'🍌',
		'🍫',
		'🍭',
		'🍕',
		'🍩',
		'🍗',
		'🍖',
		'🍬',
		'🍤',
		'🍪',
	}
	return emojies[rand.Intn(len(emojies))]
}

func hasUnicodeSupport() bool {
	return strings.Contains(os.Getenv("LANG"), "UTF-8")
}
