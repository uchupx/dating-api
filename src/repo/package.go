package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/uchupx/kajian-api/pkg/db"

	"github.com/uchupx/dating-api/src/model"
)

const (
	selectPackageQuery               = "SELECT id, name, price, description, features, status, created_at, updated_at FROM packages WHERE status = 1"
	selectPackageByIdQuery           = "SELECT id, name, price, description, features, status, created_at, updated_at FROM packages WHERE id = ?"
	selectActiveUserPackageByIdQuery = "SELECT id, user_id, feature, status FROM user_packages WHERE user_id = ? AND feature = ? AND status = 1"
	insertUserPackage                = "INSERT INTO user_packages (id, user_id, feature, status, valid_until) VALUES (?, ?, ?, ?, ?)"
)

type PackageRepo struct {
	BaseRepo
	db *db.DB
}

func (r *PackageRepo) GetPackages(ctx context.Context) ([]model.Package, error) {
	var data []model.Package

	stmt, err := r.db.FPreparexContext(ctx, selectPackageQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	rows, err := stmt.FQueryxContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var d model.Package
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Price,
			&d.Description,
			&d.Features,
			&d.Status,
			&d.CreatedAt,
			&d.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		data = append(data, d)
	}

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return data, nil
}

func (r *PackageRepo) InsertUserPackage(ctx context.Context, id string, feature string, status bool, valid *time.Time) error {
	stmt, err := r.db.FPreparexContext(ctx, insertUserPackage)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()
	_, err = stmt.FExecContext(ctx,
		r.ID(),
		id,
		feature,
		status,
		valid,
	)

	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

func (r *PackageRepo) GetPackageByID(ctx context.Context, id string) (*model.Package, error) {
	var data model.Package

	stmt, err := r.db.FPreparexContext(ctx, selectPackageByIdQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	row := stmt.FQueryRowxContext(ctx, id)
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

func (r *PackageRepo) GetActivePackageByUser(ctx context.Context, id, feature string) (*model.UserPackage, error) {
	var data model.UserPackage

	stmt, err := r.db.FPreparexContext(ctx, selectActiveUserPackageByIdQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	row := stmt.FQueryRowxContext(ctx, id, feature)
	err = row.Scan(
		&data.ID,
		&data.UserID,
		&data.Feature,
		&data.Status,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &data, nil
}

func NewPackageRepo(db *db.DB) *PackageRepo {
	return &PackageRepo{
		db: db,
	}
}
