package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shutt90/portfolio-v5/controllers"
)

func Router() *mux.Router {
	mux := mux.NewRouter()

	mux.Use(commonMiddleware)

	mux.HandleFunc("/blog", controllers.GetAllPosts).Methods("GET", "OPTIONS")
	mux.HandleFunc("/blog/create", controllers.AddPost).Methods("POST", "OPTIONS")

	return mux
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
		next.ServeHTTP(w, r)
	})
}
