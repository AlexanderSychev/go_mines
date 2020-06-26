package main

import (
	"github.com/AlexanderSychev/go_mines/Routing"
	"github.com/AlexanderSychev/go_mines/flow"
	"log"
)

const (
	AppId    = "ru.alexander-sychev.go_mines"
	AppTitle = "Go Mines"
	AppCss = "assets/styles.css"
)

func main() {
	application, err := flow.NewApplication(AppId, AppTitle, AppCss)
	if err != nil {
		log.Fatal(err)
	}

	InitRoutes()

	application.Run(func() {
		_ = flow.GetRouterInstance().RouteTo(Routing.RouteSelectGame, nil)
		if err != nil {
			log.Fatal(err)
		}
	})
}
