package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/shutt90/portfolio-v5/models"
	"github.com/shutt90/portfolio-v5/utils"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryMap := r.URL.Query()
	dbLimit := queryMap["limit"][0]

	rows, err := utils.Db.Query(context.Background(), "SELECT * FROM posts LIMIT $1", dbLimit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var p models.Post
		err := rows.Scan(p.Title, p.Body, p.Images, p.CreatedAt, p.EditedAt)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	bytes, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(bytes)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := models.Post{}

	json.NewDecoder(r.Body).Decode(&p)

	err := p.StorePost(utils.Db)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unsuccessful post request"))
		return
	}

	w.Write([]byte("saved successfully"))
	w.WriteHeader(http.StatusAccepted)
}
