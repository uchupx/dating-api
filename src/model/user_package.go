package model

import "database/sql"

type UserPackage struct {
	BaseModel
	ID        sql.NullString `db:"id"`
	Name      sql.NullString `db:"name"`
	Price     sql.NullInt64  `db:"price"`
	CreatedAt sql.NullTime   `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
	DeletedAt sql.NullTime   `db:"deleted_at"`
}

func (m *UserPackage) TableName() string {
	return "user_packages"
}
