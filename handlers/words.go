package handlers

import (
	"net/http"
	"your-words/services"
)

func HandleWord(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		services.AddWord(w, r)
	}
	if r.Method == http.MethodGet {
		services.GetAllWords(w, r)
	}
}
