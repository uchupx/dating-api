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
	findUserByUsernameEmailQuery = `
SELECT
	u.id,
	u.client_app_id,
	u.username,
	u.password,
	u.email,
	u.name,
	u.gender,
	u.address,
	u.dob,
	u.phone,
	u.created_at,
	u.updated_at,
	group_concat(DISTINCT(up.feature)) as features
FROM
	users u
	left JOIN user_packages up on u.id=up.user_id and up.status = 1
WHERE
  u.username = ? OR u.email = ?`
	findUserByIDQuery = `
SELECT
	u.id,
	u.client_app_id,
	u.username,
	u.password,
	u.email,
	u.name,
	u.gender,
	u.address,
	u.dob,
	u.phone,
	u.created_at,
	u.updated_at,
	group_concat(DISTINCT(up.feature)) as features
FROM
	users u
	left JOIN user_packages up on u.id=up.user_id and up.status = 1
WHERE
	u.id = ?
  `
	insertUserQuery     = "INSERT INTO users(id, client_app_id, username, password, email, name, gender, address, dob, phone, created_at) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	findUserRandomQuery = `
SELECT
	u.id,
	u.client_app_id,
	u.username,
	u.password,
	u.email,
	u.name,
	u.gender,
	u.address,
	u.dob,
	u.phone,
	u.created_at,
	u.updated_at,
	group_concat(up.feature) as features
FROM
	users u
	left JOIN user_packages up on u.id=up.user_id and up.status = 1
  WHERE
	u.id NOT IN (
		select
			target_user_id
		from
			reactions
		where
			user_id = u.id
			and updated_at BETWEEN ? AND ? 
	)
  AND u.id != ?
  GROUP BY u.id
  ORDER BY
	  rand()
  LIMIT 1;
`

	updateUserQuery = "UPDATE users SET name= ?, gender= ?, address = ?, dob= ?, phone =? WHERE id = ?"
)

type UserRepo struct {
	BaseRepo
	db *db.DB
}

func (r *UserRepo) FindUserByUsernameEmail(ctx context.Context, val string) (*model.User, error) {
	var user model.User

	stmt, err := r.db.FPreparexContext(ctx, findUserByUsernameEmailQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	row := stmt.FQueryRowxContext(ctx, val, val)
	err = row.Scan(
		&user.ID,
		&user.ClientAppID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Name,
		&user.Gender,
		&user.Address,
		&user.DOB,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Features,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}
	return &user, nil
}

func (r *UserRepo) FindUserRandom(ctx context.Context, selfId string, start, end time.Time) (*model.User, error) {
	var user model.User
	stmt, err := r.db.FPreparexContext(ctx, findUserRandomQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	row := stmt.FQueryRowxContext(ctx, start, end, selfId)
	err = row.Scan(
		&user.ID,
		&user.ClientAppID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Name,
		&user.Gender,
		&user.Address,
		&user.DOB,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Features,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &user, nil
}

func (r *UserRepo) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	stmt, err := r.db.FPreparexContext(ctx, findUserByIDQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	row := stmt.FQueryRowxContext(ctx, id)
	err = row.Scan(
		&user.ID,
		&user.ClientAppID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Name,
		&user.Gender,
		&user.Address,
		&user.DOB,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Features,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &user, nil
}

func (r *UserRepo) Insert(ctx context.Context, data model.User) (*string, error) {
	stmt, err := r.db.FPreparexContext(ctx, insertUserQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	id := r.ID()

	defer stmt.Close()

	_, err = stmt.FExecContext(ctx,
		id,
		data.ClientAppID.String,
		data.Username.String,
		data.Password.String,
		data.Email.String,
		data.Name.String,
		data.Gender.String,
		data.Address.String,
		data.DOB.Time,
		data.Phone.String,
		data.CreatedAt.Time,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}

	return id, nil
}

func (r *UserRepo) Update(ctx context.Context, data model.User) error {
	stmt, err := r.db.FPreparexContext(ctx, updateUserQuery)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	_, err = stmt.FExecContext(ctx,
		data.Name.String,
		data.Gender.String,
		data.Address.String,
		data.DOB.Time,
		data.Phone.String,
		data.ID.String,
	)

	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

func NewUserRepo(db *db.DB) *UserRepo {
	return &UserRepo{db: db}
}
