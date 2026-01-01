package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)
type PostStore struct {
	db *sql.DB
}
type Post struct {
	ID int64 `json:"id"`
	Content string `json:"content"`
	Title string `json:"title"`
	UserID int64 `json:"user_id"`
	Tags []string `json:"tags"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `INSERT INTO Posts (content, title, tags, user_id) VALUES ($1, $2, $3, $4) 
	RETURNING id,created_at, updated_at`
	if err := s.db.QueryRowContext(ctx, query, post.Content, post.Title,pq.Array(post.Tags), post.UserID).Scan(
		&post.ID, &post.CreatedAt, &post.UpdatedAt,
	); err != nil {
		return  err;
	}
	return nil
}