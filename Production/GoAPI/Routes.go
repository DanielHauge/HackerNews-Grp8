package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},


	Route{
		"post",
		"POST",
		"/post", PostStory,
	},
	Route{
		"latest",
		"GET",
		"/latest", GetLatest,
	},
	Route{
		"status",
		"GET",
		"/status",
		GetStatus,
	},
	Route{
		"Metrics",
		"GET",
		"/metrics",
		GetMetrics,
	},

}