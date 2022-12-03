package routes

import (
	"github.com/gorilla/mux"
	"github.com/shutt90/portfolio-v5/controllers"
)

func Router() *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/blog", controllers.GetAllPosts).Methods("GET")

	return mux
}
