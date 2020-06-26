package SelectGame

import "fmt"

// ---------------------------------------------------------------------------------------------------------------------
// "Game" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Game struct {
	width, height, mines int
	isCustom bool
}

// ---------------------------------------------------------------------------------------------------------------------
// "Game" type methods
// ---------------------------------------------------------------------------------------------------------------------

func (g Game) GetWidth() int {
	return g.width
}

func (g Game) GetHeight() int {
	return g.height
}

func (g Game) GetMines() int {
	return g.mines
}

func (g Game) IsCustom() bool {
	return g.isCustom
}

func (g Game) Label() string {
	if g.isCustom {
		return "CustomGame"
	} else {
		return fmt.Sprintf(
			"%dx%d, %d mines",
			g.width,
			g.height,
			g.mines,
		)
	}
}

func (g Game) String() string {
	return fmt.Sprintf(
		"[go_mines/SelectGame.Game](isCustom=%t, width=%d, height=%d, mines=%d)",
		g.isCustom,
		g.width,
		g.height,
		g.mines,
	)
}

// ---------------------------------------------------------------------------------------------------------------------
// "Game" type construction functions
// ---------------------------------------------------------------------------------------------------------------------

func NewGame(width, height, mines int) Game {
	return Game{
		width: width,
		height: height,
		mines: mines,
		isCustom: false,
	}
}

func NewCustomGame() Game {
	return Game{
		isCustom: true,
	}
}
