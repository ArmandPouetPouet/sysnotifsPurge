package Routes

import (
	"net/http"

	handlers "SysnotifsPurge/Handlers"

	"github.com/gorilla/mux"
)

//NewRouter router for service
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		handler = handlers.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
