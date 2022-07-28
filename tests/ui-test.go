package tests

import (
	"snake-game/components"
	"testing"
)

func TestPresenterRendersSuccessfully(t *testing.T) {
	g := components.NewGame()

	if err := g.Render(); err != nil {
		t.Fatal("Expected Game to have been rendered successfully")
	}
}
