package model

import (
	"database/sql"
)

type Reaction struct {
	BaseModel
	ID            sql.NullString `db:"id"`
	UserID        sql.NullString `db:"id"`
	TargetUserrID sql.NullString `db:"id"`
	ReactionType  sql.NullString `db:"reaction_type"`
	CreatedAt     sql.NullTime   `db:"created_at"`
	UpdatedAt     sql.NullTime   `db:"updated_at"`
}

func (m *Reaction) TableName() string {
	return "reactions"
}
