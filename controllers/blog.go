package controllers

import (
	"context"
	"net/http"

	"github.com/shutt90/portfolio-v5/utils"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	utils.Db.Query(context.Background(), "SELECT ")
}
