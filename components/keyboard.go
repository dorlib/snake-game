package components

import "github.com/nsf/termbox-go"

type keyboardEventType int

//keyboard events
const (
	MOVE keyboardEventType = 1 + iota
	RETRY
	END
)

type keyboardEvent struct {
	eventType keyboardEventType
	key       termbox.Key
}

func keyToDirection(key termbox.Key) direction {
	switch key {
	case termbox.KeyArrowLeft:
		return LEFT
	case termbox.KeyArrowRight:
		return RIGHT
	case termbox.KeyArrowUp:
		return UP
	case termbox.KeyArrowDown:
		return DOWN
	default:
		return 0
	}
}

func listenToKeyboard(evnChan chan keyboardEvent) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				evnChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowRight:
				evnChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowUp:
				evnChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowDown:
				evnChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyEsc:
				evnChan <- keyboardEvent{eventType: END, key: ev.Key}
			default:
				if ev.Ch == "r" {
					evnChan <- keyboardEvent{eventType: RETRY, key: ev.Key}
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
