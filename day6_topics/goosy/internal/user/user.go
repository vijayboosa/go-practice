package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

func (s *Store) createUserWithAuditTx(ctx context.Context, name, email string) (int64, error) {

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return 0, fmt.Errorf("begin tx: %w", err)
	}

	defer func() { _ = tx.Rollback(ctx) }()
	var userID int64
	err = tx.QueryRow(ctx, `INSERT INTO users (name, email) values ($1, $2) RETURNING ID`, name, email).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("insert user: %w", err)
	}

	b, err := json.Marshal(map[string]any{
		"source": "pgx_goosy_retry_demo",
	})

	if err != nil {
		return 0, fmt.Errorf("marshal meta: %w", err)
	}
	_, err = tx.Exec(ctx, `INSERT INTO user_audit (user_id, action, meta) VALUES($1, $2, $3::jsonb)`, userID, "USER_CREATED", b)

	if err != nil {
		return 0, fmt.Errorf("insert audit: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("commit tx: %w", err)
	}

	return userID, nil

}

func isRetryableTxErr(err error) bool {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		if pgErr.Code == "40001" || pgErr.Code == "40P01" {
			return true
		}
	}
	return false
}

func (s *Store) CreateUserWithAudit(ctx context.Context, name, email string) (int64, error) {
	maxRetries := 10
	var lastErr error
	for x := range maxRetries {
		if x >= 1 {
			delay := min(50*time.Millisecond*(1<<x), time.Second)
			select {
			case <-ctx.Done():
				return 0, ctx.Err()
			case <-time.After(delay):
			}
		}

		userId, err := s.createUserWithAuditTx(ctx, name, email)

		if err == nil {
			return userId, nil
		}

		lastErr = err
		if !isRetryableTxErr(err) {
			return 0, err
		} else {
			fmt.Printf("retry transaction for user record %s\n", name)
		}

	}

	return 0, lastErr

}
