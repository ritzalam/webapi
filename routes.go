package main

import (
	"net/http"
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
	router := mux.NewRouter()
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Create",
		"GET",
		"/api/create",
		CreateApiHandler,
	},
	Route{
		"Join",
		"GET",
		"/api/join",
		JoinApiHandler,
	},
	Route{
		"Is Meeting Running",
		"GET",
		"/api/isMeetingRunning",
		IsMeetingRunningApiHandler,
	},
	Route{
		"Get Meeting Info",
		"GET",
		"/api/getMeetingInfo",
		GetMeetingInfoApiHandler,
	},
	Route{
		"End Meeting",
		"GET",
		"/api/end",
		EndApiHandler,
	},	
	Route{
		"Get Meetings",
		"GET",
		"/api/getMeetings",
		GetMeetingsApiHandler,
	},	
	Route{
		"Get Default Config XML",
		"GET",
		"/api/getDefaultConfigXML",
		GetDefaultConfigXMLApiHandler,
	},
	Route{
		"Set Config XML",
		"GET",
		"/api/setConfigXML",
		SetConfigXMLApiHandler,
	},
	Route{
		"Enter",
		"GET",
		"/api/enter",
		EnterApiHandler,
	},
	Route{
		"Config XML",
		"GET",
		"/api/configXML",
		ConfigXMLApiHandler,
	},
	Route{
		"Sign Out",
		"GET",
		"/api/signOut",
		SignOutApiHandler,
	},
	Route{
		"Get Recordings",
		"GET",
		"/api/getRecordings",
		GetRecordingsApiHandler,
	},
	Route{
		"Publish Recordings",
		"GET",
		"/api/publishRecordings",
		PublishRecordingsApiHandler,
	},
	Route{
		"Delete Recordings",
		"GET",
		"/api/deleteRecordings",
		DeleteRecordingsApiHandler,
	},
}
