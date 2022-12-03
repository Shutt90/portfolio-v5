package utils

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

const (
	ErrConnToDb = "could not connect to database"
)

var Db *pgx.Conn

func DbConnect() {
	dsn := os.Getenv("DB_CONNECT")
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal("could not connect to db: ", err)
	}
	defer conn.Close(context.Background())
	log.Println("db connection established")

	Db = conn
}
