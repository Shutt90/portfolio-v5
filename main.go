package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/shutt90/portfolio-v5/routes"
	"github.com/shutt90/portfolio-v5/utils"
)

func main() {
	godotenv.Load()
	utils.DbConnect()

	path := filepath.Join("sql", "blog.sql")
	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		log.Fatal("path not created with sql table")
	}

	sql := string(c)

	_, err := utils.Db.Exec(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
		log.Fatal("could not create table")
	}

	mux := routes.Router()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
