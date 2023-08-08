package app

import (
	"sprint/core"
)

// Controller Declaration
type AppController struct {
	Name string
	Path string
}

func (reflect AppController) Init() (string, string) {
	return reflect.Name, reflect.Path
}

// Controller Routes
func (reflect AppController) Routes() []core.Route {
	return []core.Route{
		{
			Endpoint: "/Hello/:name",
			Method:   "GET",
			Function: reflect.Hello,
		},
		{
			Endpoint: "/Test",
			Method:   "POST",
			Function: reflect.Test,
		},
	}
}

// Controller Methods
func (reflect AppController) Hello(request core.Request) core.Response {
	return core.Response{Content: "Hello World"}
}

func (reflect AppController) Manger(request core.Request) core.Response {
	return core.Response{Content: "Hello World"}
}

func (reflect AppController) Test(request core.Request) core.Response {
	return core.Response{Content: "Hello World"}
}
