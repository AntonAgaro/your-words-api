package handlers

import (
	"net/http"
	"your-words/services"
)

func HandleTopic(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		services.GetTopics(w, r)
	}
}
