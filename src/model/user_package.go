package model

import "database/sql"

type UserPackage struct {
	BaseModel
	ID         sql.NullString `db:"id"`
	UserID     sql.NullString `db:"user_id"`
	Feature    sql.NullString `db:"feature"`
	Status     sql.NullBool   `db:"status"`
	ValidUntil sql.NullTime   `db:"valid_until"`
	CreatedAt  sql.NullTime   `db:"created_at"`
}

func (m *UserPackage) TableName() string {
	return "user_packages"
}
