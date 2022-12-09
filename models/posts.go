package models

import (
	"context"
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
	_, err := db.Exec(context.Background(), "INSERT INTO post ($1, $2, $3, $4, $5)", &p.Title, &p.Body, &p.Images, &p.CreatedAt, &p.EditedAt)
	if err != nil {
		return err
	}

	return nil
}
