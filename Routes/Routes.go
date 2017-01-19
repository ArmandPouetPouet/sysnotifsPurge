package Routes

import (
	"net/http"

	handlers "SysnotifsPurge/Handlers"
)

//Route a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes route list
type Routes []Route

var routes = Routes{
	Route{
		Name:        "CheckRatio",
		Method:      "GET",
		Pattern:     "/CheckRatio",
		HandlerFunc: handlers.CheckRatio,
	},
	Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/index",
		HandlerFunc: handlers.Index,
	},
}
