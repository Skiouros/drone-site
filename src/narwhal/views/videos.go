package views

import "net/http"
import "narwhal/server"
import "narwhal/database"

func videoGetAll(w http.ResponseWriter, r *http.Request) {
	ServeJson(w, database.GetAllVideos())
}

func videoNew(w http.ResponseWriter, r *http.Request) {
	msg := Message{}

	pos := r.FormValue("pos")
	if pos == "" {
		msg.Errors = []string{ "Must select a location" }
		ServeJson(w, msg)
		return
	}

	url := r.FormValue("video")
	if url == "" {
		msg.Errors = []string{ "Must enter a video url" }
		ServeJson(w, msg)
		return
	}

	database.CreateVideo(pos, url)

	msg.Content = "ok"
	ServeJson(w, msg)
}

func init() {
	server := server.GetServer()

	server.Route("/videos", videoNew).Methods("POST")
	server.Route("/videos", videoGetAll).Methods("GET")
}
