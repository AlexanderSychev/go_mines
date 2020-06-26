package SelectGame

import (
	"fmt"
	"github.com/AlexanderSychev/go_mines/flow"
	"log"
	"reflect"
	"strings"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
// "Model" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Model [4]Game

// ---------------------------------------------------------------------------------------------------------------------
// "Model" type methods
// ---------------------------------------------------------------------------------------------------------------------

func (m *Model) GetGamesCount() int {
	return len(m)
}

func (m *Model) GetAllGames() []Game {
	return m[:]
}

func (m *Model) GetGame(index int) Game {
	if index < 0 {
		index = 0
	}
	if index > m.GetGamesCount()-1 {
		index = m.GetGamesCount()-1
	}
	return m[index]
}

func (m *Model) String() string {
	strList := make([]string, 0)
	for _, game := range m {
		strList = append(strList, game.String())
	}
	return fmt.Sprintf("[go_mines/SelectGame.Model][%s]", strings.Join(strList, ", "))
}

func (m *Model) HandleUnknownMessage(message flow.Message) ([]reflect.Value, error) {
	log.Printf(
		"[%s] %v Does not understand: %v",
		time.Now().Format(time.UnixDate),
		m,
		message,
	)
	return nil, nil
}

// ---------------------------------------------------------------------------------------------------------------------
// "Model" type construction function
// ---------------------------------------------------------------------------------------------------------------------

func NewModel() *Model {
	return &Model{
		NewGame(8, 8, 10),
		NewGame(16, 16, 40),
		NewGame(30, 16, 99),
		NewCustomGame(),
	}
}
