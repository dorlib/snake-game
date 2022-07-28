package tests

import (
	"github.com/nsf/termbox-go"
	"snake-game/components"
	"testing"
)

func testKeyToDirectionDefault(t *testing.T) {
	d := components.KeyToDirection(termbox.KeyEsc)
	if d != 0 {
		t.Fatalf("Expected direction to be 0 but go %v", d)
	}
}

func testKeyToDirectionRight(t *testing.T) {
	d := components.KeyToDirection(termbox.KeyArrowRight)
	if d != RIGHT {
		t.Fatalf("expected direction to be RIGHT but got %v", d)
	}
}

func testKeyToDirectionLeft(t *testing.T) {
	d := components.KeyToDirection(termbox.KeyArrowLeft)
	if d != LEFT {
		t.Fatalf("expected direction to be LEFT but got %v", d)
	}
}

func testKeyToDirectionUp(t *testing.T) {
	d := components.KeyToDirection(termbox.KeyArrowUp)
	if d != UP {
		t.Fatalf("expected direction to be UP but got %v", d)
	}
}

func testKeyToDirectionDwn(t *testing.T) {
	d := components.KeyToDirection(termbox.KeyArrowDown)
	if d != DOWN {
		t.Fatalf("expected direction to be DOWN but got %v", d)
	}
}
