package model

import "database/sql"

type User struct {
	BaseModel
	ID          sql.NullString `db:"id"`
	ClientAppID sql.NullString `db:"client_app_id"`
	Username    sql.NullString `db:"username"`
	// Name      sql.NullString `db:"name"`
	Password  sql.NullString `db:"password"`
	Email     sql.NullString `db:"email"`
	CreatedAt sql.NullTime   `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
}

func (m *User) TableName() string {
	return "users"
}
