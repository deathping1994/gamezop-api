package main
import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
		"CarByID",
		[]string{"GET","PATCH","DELETE", "OPTIONS"},
		"/cars/{id}/",
		CarByID,
	},
	Route{
		"CarApi",
		[]string{"GET","POST", "OPTIONS"},
		"/cars/",
		CarApi,
	},
}

func withCORS(fn http.HandlerFunc) http.HandlerFunc {
 return func(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Access-Control-Allow-Origin", "*")
   w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, DELETE")
   fn(w, r)
 }
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = withCORS(route.HandlerFunc)
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}