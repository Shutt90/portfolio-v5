package models

import "github.com/jackc/pgx"

type Post struct {
	Title     string  `json:"title,omitempty"`
	Body      string  `json:"body,omitempty"`
	Images    []Image `json:"image,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
	EditedAt  string  `json:"edited_at,omitempty"`
}

type Image struct {
	Title       string `json:"title,omitempty"`
	Url         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

func (p Post) StorePost(db *pgx.Conn) (bool, error) {
	_, err := db.Exec("INSERT INTO post ($1, $2, $3, $4, $5)", p.Title, p.Body, p.Images, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return false, err
	}

	return true, nil
}
