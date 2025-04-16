package main

import (
	"log"
	"net/http"

	"github.com/yourusername/comment_moderation_service/internal/db"
	"github.com/yourusername/comment_moderation_service/internal/handler"
)

func main() {
	db.InitDB()
	defer db.DB.Close()

	http.HandleFunc("/comments", handler.GetAllComments)
	http.HandleFunc("/comments/", handler.GetComment)
	http.HandleFunc("/comments", handler.CreateComment)
	http.HandleFunc("/comments/", handler.UpdateComment)
	http.HandleFunc("/comments/", handler.DeleteComment)

	port := ":8080"
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
