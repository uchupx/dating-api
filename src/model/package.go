package model

import "database/sql"

type Package struct {
	BaseModel
	ID          sql.NullString `db:"id"`
	Name        sql.NullString `db:"name"`
	Price       sql.NullInt64  `db:"price"`
	Description sql.NullString `db:"description"`
	Features    sql.NullString `db:"features"`
	Status      sql.NullBool   `db:"status"`
	CreatedAt   sql.NullTime   `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at"`
}

func (m *Package) TableName() string {
	return "packages"
}
