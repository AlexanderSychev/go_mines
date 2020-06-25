package main

import (
	"github.com/AlexanderSychev/go_mines/CustomGame"
	"github.com/AlexanderSychev/go_mines/flow"
	"log"
)

const (
	AppId    = "ru.alexander-sychev.go_mines"
	AppTitle = "Go Mines"
	AppCss = "assets/styles.css"
	RouteCustomGame = "custom-game"
)

func main() {
	application, err := flow.NewApplication(AppId, AppTitle, AppCss)
	if err != nil {
		log.Fatal(err)
	}

	router := flow.GetRouterInstance()
	router.AddRoute(RouteCustomGame, CustomGame.PageCreator)
	application.Run(func() {
		_ = router.RouteTo(RouteCustomGame, nil)
		if err != nil {
			log.Fatal(err)
		}
	})
}
