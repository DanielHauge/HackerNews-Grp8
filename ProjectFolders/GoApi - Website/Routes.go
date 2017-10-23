package main

import "net/http"

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
		"CreateUser",
		"POST",
		"/create",
		CreateUser,
	},
	Route{
		"VerifyUser",
		"POST",
		"/login",
		Login,
	},
	Route{
		"GetLatestStory",
		"POST",
		"/stories",
		GetLatestStories,
	},
	Route{
		"GetStoryByID",
		"GET",
		"/stories/{storyid}",
		GetStoryByID,
	},
	Route{
		"GetComment",
		"GET",
		"/comments/{storyid}",
		GetComments,
	},
}