package Game

import (
	"github.com/AlexanderSychev/go_mines/flow"
	"log"
	"math/rand"
	"reflect"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------------------------------------------------

const (
	tickerInterval = time.Second
	tickerStop = time.Minute * 20
)

// ---------------------------------------------------------------------------------------------------------------------
// "Cell" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Cell struct {
	IsOpened, IsMarked, HasMine bool
	MinesAround int
}

// ---------------------------------------------------------------------------------------------------------------------
// "Model" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Model struct {
	ticker *time.Ticker
	field [][]Cell
}

// ---------------------------------------------------------------------------------------------------------------------
// "Model" type private methods
// ---------------------------------------------------------------------------------------------------------------------

func (m *Model) runTicker() {
	go func() {
		startTime := time.Now()
		go func() {
			for t := range m.ticker.C {
				diff := startTime.Sub(t)
				message := flow.NewMessageWithArgs("UpdateTime", diff)
				_, err := flow.GetBrokerInstance().SendToActor(flow.ControllerActor, message)
				if err != nil {
					log.Println(err)
				}
			}
		}()
		time.Sleep(tickerStop)
		_, err := flow.GetBrokerInstance().SendToActor(flow.ControllerActor, flow.NewMessage("Timeout"))
		if err != nil {
			log.Println(err)
		}
	}()
}

func (m *Model) generateMines(mines int) {
	rand.Seed(time.Now().UTC().UnixNano())
	width := m.width()
	height := m.height()

	for i := 0; i < mines; i++ {
		for {
			x := rand.Intn(height)
			y := rand.Intn(width)
			if !m.field[x][y].HasMine {
				m.field[x][y].HasMine = true
				break
			}
		}
	}
}

func (m *Model) getNeighbors(x, y int) []Point {
	width := m.width()
	height := m.height()

	result := make([]Point, 0)

	// Topper line (lesser x)
	if x > 0 {
		if y > 0 {
			result = append(result, NewPoint(x-1, y-1))
		}
		result = append(result, NewPoint(x-1, y))
		if y < width-1 {
			result = append(result, NewPoint(x-1, y+1))
		}
	}

	// Same line (same x)
	if y > 0 {
		result = append(result, NewPoint(x, y-1))
	}
	if y < width-1 {
		result = append(result, NewPoint(x, y+1))
	}

	// Downer line
	if x < height-1 {
		if y > 0 {
			result = append(result, NewPoint(x+1, y-1))
		}
		result = append(result, NewPoint(x+1, y))
		if y < width-1 {
			result = append(result, NewPoint(x+1, y+1))
		}
	}

	return result
}

func (m *Model) getMinesAround(x, y int) int {
	result := 0

	neighbors := m.getNeighbors(x, y)

	for _, point := range neighbors {
		if m.field[point.X()][point.Y()].HasMine {
			result++
		}
	}

	return result
}

func (m *Model) calculateMinesAround() {
	for x, line := range m.field {
		for y, _ := range line {
			m.field[x][y].MinesAround = m.getMinesAround(x, y)
		}
	}
}

func (m *Model) height() int {
	return len(m.field)
}

func (m *Model) width() int {
	return len(m.field[0])
}

// ---------------------------------------------------------------------------------------------------------------------
// "Model" type public methods
// ---------------------------------------------------------------------------------------------------------------------

func (m *Model) GetEmptyNotMarkedNeighbors(x, y int) []Point {
	result := make([]Point, 0)

	neighbors := m.getNeighbors(x, y)

	for _, point := range neighbors {
		if !m.field[point.X()][point.Y()].HasMine && !m.field[point.X()][point.Y()].IsMarked {
			result = append(result, point)
		}
	}

	return result
}

func (m *Model) GetAllWithMines() []Point {
	result := make([]Point, 0)

	for x, line := range m.field {
		for y, cell := range line {
			if cell.HasMine {
				result = append(result, NewPoint(x, y))
			}
		}
	}

	return result
}

func (m *Model) IsAllCellsOpened() bool {
	result := true

	for _, line := range m.field {
		for _, cell := range line {
			result = result && cell.IsOpened
		}
	}

	return result
}

func (m *Model) ToggleMark(x, y int) bool {
	m.field[x][y].IsMarked = !m.field[x][y].IsMarked
	return m.field[x][y].IsMarked
}

func (m *Model) Open(x, y int) (bool, int) {
	m.field[x][y].IsMarked = false
	m.field[x][y].IsOpened = true
	return m.field[x][y].HasMine, m.field[x][y].MinesAround
}

func (m *Model) HandleUnknownMessage(message flow.Message) ([]reflect.Value, error) {
	log.Printf(
		"[%s] [go_mines/Game.Model]() Does not understand: %v",
		time.Now().Format(time.UnixDate),
		message,
	)
	return nil, nil
}

// ---------------------------------------------------------------------------------------------------------------------
// "Model" type construction
// ---------------------------------------------------------------------------------------------------------------------

func NewModel(width, height, mines int) *Model {
	field := make([][]Cell, height, height)
	for x := 0; x < height; x++ {
		field[x] = make([]Cell, width, width)
		for y := 0; y < width; y++ {
			field[x][y] = Cell{}
		}
	}

	ticker := time.NewTicker(tickerInterval)

	result := &Model{
		field: field,
		ticker: ticker,
	}
	result.generateMines(mines)
	result.calculateMinesAround()
	result.runTicker()

	return result
}
