/*
 * star world API
 *
 * the api to query the information about *Star War* you can check all the Star Wars data you've ever wanted Planets Spaceships Vehicles People Films and Species From all SEVEN Star Wars films
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {

}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"FilmsGet",
		strings.ToUpper("Get"),
		"/films",
		FilmsGet,
	},

	Route{
		"FilmsIdGet",
		strings.ToUpper("Get"),
		"/films/{id}",
		FilmsIdGet,
	},

	Route{
		"PeopleGet",
		strings.ToUpper("Get"),
		"/people",
		PeopleGet,
	},

	Route{
		"PeopleIdGet",
		strings.ToUpper("Get"),
		"/people/{id}",
		PeopleIdGet,
	},

	Route{
		"PlanetsGet",
		strings.ToUpper("Get"),
		"/planets",
		PlanetsGet,
	},

	Route{
		"PlanetsIdGet",
		strings.ToUpper("Get"),
		"/planets/{id}",
		PlanetsIdGet,
	},

	Route{
		"SpeciesGet",
		strings.ToUpper("Get"),
		"/species",
		SpeciesGet,
	},

	Route{
		"SpeciesIdGet",
		strings.ToUpper("Get"),
		"/species/{id}",
		SpeciesIdGet,
	},

	Route{
		"StarshipsGet",
		strings.ToUpper("Get"),
		"/starships",
		StarshipsGet,
	},

	Route{
		"StarshipsIdGet",
		strings.ToUpper("Get"),
		"/starships/{id}",
		StarshipsIdGet,
	},

	Route{
		"UserLogoutPost",
		strings.ToUpper("Post"),
		"/user/logout",
		UserLogoutPost,
	},

	Route{
		"UserSignInPost",
		strings.ToUpper("Post"),
		"/user/signIn",
		UserSignInPost,
	},

	Route{
		"UserSignUpPost",
		strings.ToUpper("Post"),
		"/user/signUp",
		UserSignUpPost,
	},

	Route{
		"UserUpdatePost",
		strings.ToUpper("Post"),
		"/user/update",
		UserUpdatePost,
	},

	Route{
		"VehiclesGet",
		strings.ToUpper("Get"),
		"/vehicles",
		VehiclesGet,
	},

	Route{
		"VehiclesIdGet",
		strings.ToUpper("Get"),
		"/vehicles/{id}",
		VehiclesIdGet,
	},
}
