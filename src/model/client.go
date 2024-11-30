package model

import (
	"database/sql"
)

type Client struct {
	BaseModel
	ID        sql.NullString `db:"id"`
	Name      sql.NullString `db:"name"`
	Key       sql.NullString `db:"key"`
	Secret    sql.NullString `db:"secret"`
	CreatedAt sql.NullTime   `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
	RevokedAt sql.NullTime   `db:"revoked_at"`
}

func (m *Client) TableName() string {
	return "client_apps"
}
