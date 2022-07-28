package tests

import (
	"os"
	"snake-game/components"
	_ "snake-game/components"
	"testing"
)

func testFoodDefaultPoint(t *testing.T) {
	food := components.NewFood(10, 10)
	if food.Score != 10 {
		t.Fatalf("Expected food dafault score to be 10 but go %v", food.Score)
	}
}

func testEmoji(t *testing.T) {
	food := components.NewFood(10, 10)
	if string(food.Emoji) == "" {
		t.Fatalf("error: got blank food emoji")
	}
}

func testFoodFallback(t *testing.T) {
	os.Setenv("LANG", "c")
	food := components.NewFood(10, 10)
	if string(food.Emoji) != "@" {
		t.Fatalf("food emoji expected to be @")
	}
}
