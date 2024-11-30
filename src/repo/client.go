package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uchupx/dating-api/src/model"
	"github.com/uchupx/kajian-api/pkg/db"
)

const (
	findAppsByKeyQuery = "SELECT id,name,`key`,secret  FROM client_apps WHERE `key` = ?"
	insertClientQuery  = "INSERT INTO client_apps(id, name, `key`, secret) VALUES (?, ?, ?, ?)"
)

type ClientRepo struct {
	BaseRepo
	db *db.DB
}

func (r *ClientRepo) FindAppsByKey(ctx context.Context, val string) (*model.Client, error) {
	var data model.Client

	stmt, err := r.db.FPreparexContext(ctx, findAppsByKeyQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	row := stmt.FQueryRowxContext(ctx, val)
	err = row.Scan(&data.ID, &data.Name, &data.Key, &data.Secret)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &data, nil
}

func (r *ClientRepo) Insert(ctx context.Context, data model.Client) (*string, error) {
	stmt, err := r.db.FPreparexContext(ctx, insertClientQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	id := r.ID()

	defer stmt.Close()

	_, err = stmt.FExecContext(ctx,
		id,
		data.Name.String,
		data.Key.String,
		data.Secret.String,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}

	return id, nil
}

func NewClientRepo(db *db.DB) *ClientRepo {
	return &ClientRepo{
		db: db,
	}
}
