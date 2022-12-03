package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/shutt90/portfolio-v5/routes"
	"github.com/shutt90/portfolio-v5/utils"
)

func main() {
	godotenv.Load()
	utils.DbConnect()

	mux := routes.Router()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
