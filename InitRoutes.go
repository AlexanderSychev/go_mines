package main

import (
	"github.com/AlexanderSychev/go_mines/CustomGame"
	"github.com/AlexanderSychev/go_mines/Game"
	"github.com/AlexanderSychev/go_mines/Routing"
	"github.com/AlexanderSychev/go_mines/SelectGame"
	"github.com/AlexanderSychev/go_mines/flow"
)

func InitRoutes() {
	router := flow.GetRouterInstance()
	router.AddRoute(Routing.RouteCustomGame, CustomGame.PageCreator)
	router.AddRoute(Routing.RouteSelectGame, SelectGame.PageCreator)
	router.AddRoute(Routing.RouteGame, Game.PageCreator)
}
