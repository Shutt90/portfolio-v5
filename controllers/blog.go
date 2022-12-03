package controllers

import (
	"context"
	"net/http"

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

	rows.Scan()

}
