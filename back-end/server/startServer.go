package server

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Server() {
	var err error
	db, err = sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/createpost", createPostHandler)
	http.HandleFunc("/createcomment", createCommentHandler)
	http.HandleFunc("/likepost", likePostHandler)
	http.HandleFunc("/dislikepost", dislikePostHandler)
	http.HandleFunc("/filterposts", filterPostsHandler)
	log.Fatal(http.ListenAndServe(":0666", nil))
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

}

func loginHandler(w http.ResponseWriter, r *http.Request) {

}

func createPostHandler(w http.ResponseWriter, r *http.Request) {

}

func createCommentHandler(w http.ResponseWriter, r *http.Request) {

}

func likePostHandler(w http.ResponseWriter, r *http.Request) {

}

func dislikePostHandler(w http.ResponseWriter, r *http.Request) {

}

func filterPostsHandler(w http.ResponseWriter, r *http.Request) {

}
