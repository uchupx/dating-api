package model

import "database/sql"

type User struct {
	BaseModel
	ID          sql.NullString `db:"id"`
	ClientAppID sql.NullString `db:"client_app_id"`
	Username    sql.NullString `db:"username"`
	Name        sql.NullString `db:"name"`
	Gender      sql.NullString `db:"gender"`
	Address     sql.NullString `db:"address"`
	Phone       sql.NullString `db:"phone"`
	Password    sql.NullString `db:"password"`
	Email       sql.NullString `db:"email"`
	DOB         sql.NullTime   `db:"dob"`
	CreatedAt   sql.NullTime   `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`

	Features sql.NullString `db:"features"`
}

func (m *User) TableName() string {
	return "users"
}
