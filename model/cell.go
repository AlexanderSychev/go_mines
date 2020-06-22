package model

// -----------------------------------------------------------------------------
// Constants
// -----------------------------------------------------------------------------

const (
	mineString    = "*"
	markedString  = "x"
	defaultString = "_"
)

// -----------------------------------------------------------------------------
// "Cell" type definition
// -----------------------------------------------------------------------------

type Cell struct {
	hasMine  bool
	isMarked bool
	isOpened bool
}

// Methods

func (cell *Cell) ToggleMark() bool {
	canToggle := !cell.isOpened
	if canToggle {
		cell.isMarked = !cell.isMarked
	}
	return canToggle
}

func (cell *Cell) Open() {
	cell.isMarked = false
	cell.isOpened = true
}

func (cell Cell) HasMine() bool {
	return cell.hasMine
}

func (cell Cell) IsOpened() bool {
	return cell.isOpened
}

func (cell Cell) IsMarked() bool {
	return cell.isMarked
}

func (cell Cell) String() string {
	if cell.isOpened && cell.hasMine {
		return mineString
	} else if cell.isMarked {
		return markedString
	}

	return defaultString
}

// Constructor

func NewCell(hasMine bool) *Cell {
	return &Cell{hasMine, false, false}
}
