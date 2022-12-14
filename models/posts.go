package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Post struct {
	Title     string    `json:"title,omitempty"`
	Body      string    `json:"body,omitempty"`
	Images    []Image   `json:"image,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	EditedAt  time.Time `json:"edited_at,omitempty"`
}

type Image struct {
	Title       string `json:"title,omitempty"`
	Url         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

func (p Post) StorePost(db *pgx.Conn) error {
	fmt.Println(p.Title, p.Body, p.Images, p.CreatedAt)
	_, err := db.Exec(context.Background(), "INSERT INTO blog (title, body, created_at) VALUES ($1, $2, $3)", p.Title, p.Body, p.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
