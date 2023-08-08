package app

import (
	"sprint/core"
)

// Module Definition
var AppModule = core.Module{
	Name: "AppModule",
	Controllers: []core.Controller{
		AppController{Path: "app", Name: "AppController"},
	},
}
