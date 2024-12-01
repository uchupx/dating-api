package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uchupx/dating-api/src/model"
	"github.com/uchupx/kajian-api/pkg/db"
)

const (
	selectPackageQuery = "SELECT id, name, price, description, features, status, created_at, updated_at FROM packages WHERE status = 1"
	insertUserPackage  = "INSERT INTO user_packages (id, user_id, feature, status, valid_until) VALUES (?, ?, ?, ?, ?)"
)

type PackageRepo struct {
	BaseRepo
	db *db.DB
}

func (r *PackageRepo) GetPackages(ctx context.Context) (*model.Package, error) {
	var data model.Package

	stmt, err := r.db.FPreparexContext(ctx, selectPackageQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	row := stmt.FQueryRowxContext(ctx)
	err = row.Scan(
		&data.ID,
		&data.Name,
		&data.Price,
		&data.Description,
		&data.Features,
		&data.Status,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &data, nil
}

func (r *PackageRepo) InsertUserPackage(ctx context.Context, data *model.UserPackage) error {
	stmt, err := r.db.FPreparexContext(ctx, insertUserPackage)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()
	_, err = stmt.FExecContext(ctx,
		data.ID,
		data.UserID,
		data.Feature,
		data.Status,
		data.ValidUntil,
	)

	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

func NewPackageRepo(db *db.DB) *PackageRepo {
	return &PackageRepo{
		db: db,
	}
}
