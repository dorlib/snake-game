package components

import (
	"math/rand"
	"os"
	"strings"
)

type Food struct {
	Emoji rune
	X     int
	Y     int
	Score int
}

func NewFood(x int, y int) *Food {
	return &Food{
		Emoji: getFoodEmoji(),
		X:     x,
		Y:     y,
		Score: 10,
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
