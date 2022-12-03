package main

import (
	"github.com/joho/godotenv"
	"github.com/shutt90/portfolio-v5/utils"
)

func main() {
	godotenv.Load()
	utils.DbConnect()
}
