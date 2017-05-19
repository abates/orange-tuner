package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"ChannelsGet",
		"GET",
		"/channels",
		ChannelsGet,
	},

	Route{
		"ChannelsIdGet",
		"GET",
		"/channels/{id}",
		ChannelsIdGet,
	},

	Route{
		"EpgChannelIdGet",
		"GET",
		"/epg/{channel_id}",
		EpgChannelIdGet,
	},

	Route{
		"EpgGet",
		"GET",
		"/epg",
		EpgGet,
	},

	Route{
		"ScanGet",
		"GET",
		"/scan",
		ScanGet,
	},

	Route{
		"ScanPost",
		"POST",
		"/scan",
		ScanPost,
	},

	Route{
		"StreamChannelIdPlaylistM3u8Delete",
		"DELETE",
		"/stream/{channel_id}/playlist.m3u8",
		StreamChannelIdPlaylistM3u8Delete,
	},

	Route{
		"StreamChannelIdPlaylistM3u8Get",
		"GET",
		"/stream/{channel_id}/playlist.m3u8",
		StreamChannelIdPlaylistM3u8Get,
	},

	Route{
		"StreamChannelIdSegmentNameGet",
		"GET",
		"/stream/{channel_id}/{segment_name}",
		StreamChannelIdSegmentNameGet,
	},
}
