package user

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt time.Time
}

type Store struct {
	pool *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{pool: pool}
}

func (s *Store) List(ctx context.Context, limit int) ([]User, error) {
	rows, err := s.pool.Query(ctx, `SELECT id, name, email, created_at FROM users ORDER BY id LIMIT $1`, limit)
	if err != nil {
		return nil, fmt.Errorf("query users: %w", err)
	}

	defer rows.Close()

	out := make([]User, 0, limit)

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}
		out = append(out, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return out, nil

}

func (s *Store) Create(ctx context.Context, name, email string) (int64, error) {
	var id int64

	err := s.pool.QueryRow(ctx, `INSERT INTO users (name, email) values ($1, $2) RETURNING id`, name, email).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("insert user: %w", err)
	}

	return id, nil

}
