package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/uchupx/dating-api/src/model"
	"github.com/uchupx/kajian-api/pkg/db"
)

const (
	findRefreshTokenQuery = "SELECT id, user_id, client_app_id, expired_at FROM refresh_tokens WHERE token = ?"
	insertRefreshToken    = "INSERT INTO refresh_tokens(id, user_id, client_app_id, token, expired_at) VALUES (?, ?, ?, ?, ?)"
)

type RefreshTokenRepo struct {
	BaseRepo
	db *db.DB
}

func (r *RefreshTokenRepo) FindAppsByKey(ctx context.Context, val string) (*model.RefreshToken, error) {
	var data model.RefreshToken

	stmt, err := r.db.FPreparexContext(ctx, findAppsByKeyQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	row := stmt.FQueryRowxContext(ctx, val)
	err = row.Scan(&data.ID, &data.UserID, &data.ClientAppID, &data.ExpiredAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &data, nil
}

func (r *RefreshTokenRepo) Insert(ctx context.Context, userId string, clientAppId string, token string, expiredAt time.Time) (*string, error) {
	stmt, err := r.db.FPreparexContext(ctx, insertRefreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	id := r.ID()

	defer stmt.Close()

	_, err = stmt.FExecContext(ctx,
		id,
		userId,
		clientAppId,
		token,
		expiredAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}

	return id, nil
}

func NewRefreshTokenRepo(db *db.DB) *RefreshTokenRepo {
	return &RefreshTokenRepo{
		db: db,
	}
}
