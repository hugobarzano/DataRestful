package controller

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"datacore/mongo"
)


var controller = Controller{Storer: mongo.Storer{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes {
	//Route {
	//	"Authentication",
	//	"POST",
	//	"/get-token",
	//	controller.GetToken,
	//},
	Route {
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route {
		"AddDataset",
		"POST",
		"/AddDataset",
		controller.AddDataset,
	},
	Route {
		"UpdateDataset",
		"PUT",
		"/UpdateDataset",
		controller.UpdateDataset,
	},
	// Get Dataset by {id}
	Route {
		"GetDataset",
		"GET",
		"/datasets/{id}",
		controller.GetDataset,
	},
	// Delete Dataset by {id}
	Route {
		"DeleteDataset",
		"DELETE",
		"/deleteDataset/{id}",
		controller.DeleteDataset,
	},
	// Search Dataset with string
	Route {
		"SearchDataset",
		"GET",
		"/Search/{query}",
		controller.SearchDataset,
	},
	Route {
		"PerformsOperation",
		"POST",
		"/Operation/",
		controller.PerformsOperation,
	},	Route {
		"RegisterServices",
		"POST",
		"/Services/",
		controller.AddService,
	},	Route {
		"ListServices",
		"GET",
		"/Services/",
		controller.ListServices,
	},	Route {
		"DeleteService",
		"DELETE",
		"/RemoveService/{id}",
		controller.DeleteService,
	}}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

