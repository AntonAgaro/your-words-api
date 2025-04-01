package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"your-words/database"
	"your-words/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//handlers.AddWord()
	database.ConnectDb(os.Getenv("DATABASE_URL"))

	http.HandleFunc("/words", handlers.HandleWord)
	http.HandleFunc("/topics", handlers.HandleTopic)

	err = http.ListenAndServe(":8080", enableCORS(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}

}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
