package model

import (
	"database/sql"
)

type RefreshToken struct {
	BaseModel
	ID          sql.NullString `db:"id"`
	UserID      sql.NullString `db:"user_id"`
	ClientAppID sql.NullString `db:"client_app_id"`
	ExpiredAt   sql.NullTime   `db:"expired_at"`
}

func (m *RefreshToken) TableName() string {
	return "refresh_tokens"
}
