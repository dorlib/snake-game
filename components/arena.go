package components

type arena struct {
	height int,
	width int,
	hasFood func(*arena, cordiante) bool,
	snake *snake,
	food *food,
	scoreChan chan (int)
}


