package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uchupx/dating-api/src/model"
	"github.com/uchupx/kajian-api/pkg/db"
)

const (
	selectReactionQuery = "SELECT id, user_id, target_user_id, reaction_type, created_at, updated_at FROM reactions WHERE user_id = ? and target_user_id = ?"
	insertReactionQuery = "INSERT INTO reactions(id, user_id, target_user_id, reaction_type) VALUES (?, ?, ?, ?)"
	updateReactionQuery = "UPDATE reactions SET reaction_type = ? WHERE id = ?"
)

type ReactionRepo struct {
	BaseRepo
	db *db.DB
}

func (r *ReactionRepo) FindByUserIdTargetIdPair(ctx context.Context, userId, targetId string) (*model.Reaction, error) {
	var reaction model.Reaction

	stmt, err := r.db.FPreparexContext(ctx, selectReactionQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	row := stmt.FQueryRowxContext(ctx, userId, targetId)
	err = row.Scan(
		&reaction.ID,
		&reaction.UserID,
		&reaction.TargetUserID,
		&reaction.ReactionType,
		&reaction.CreatedAt,
		&reaction.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &reaction, nil
}

func (r *ReactionRepo) Insert(ctx context.Context, userId, targetId string, reactType int8) (*string, error) {
	stmt, err := r.db.FPreparexContext(ctx, insertReactionQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	id := r.ID()

	defer stmt.Close()

	_, err = stmt.FExecContext(ctx,
		id,
		userId,
		targetId,
		reactType,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}

	return id, nil
}

func (r *ReactionRepo) Update(ctx context.Context, reactType int8, id string) error {
	stmt, err := r.db.FPreparexContext(ctx, updateReactionQuery)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	_, err = stmt.FExecContext(ctx,
		reactType,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

func NewReactionRepo(db *db.DB) *ReactionRepo {
	return &ReactionRepo{
		db: db,
	}
}
