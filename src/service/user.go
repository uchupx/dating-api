package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/uchupx/dating-api/pkg/database/redis"
	"github.com/uchupx/dating-api/pkg/helper"
	"github.com/uchupx/dating-api/src/dto"
	"github.com/uchupx/dating-api/src/repo"
	"github.com/uchupx/kajian-api/pkg/errors"
)

type UserService struct {
	UserRepo     *repo.UserRepo
	ReactionRepo *repo.ReactionRepo

	Redis *redis.Redis
}

func (UserService) name() string {
	return "UserService"
}

func (s *UserService) findByID(ctx context.Context, id string) (*dto.User, error) {
	model, err := s.UserRepo.FindUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s - findUserByID] error when find user by id: %w", s.name(), err)
	} else if model == nil {
		return nil, errors.ErrNotFound
	}

	var user dto.User
	user.Model(model)

	return &user, nil
}

func (s *UserService) FindUserByID(ctx context.Context, id string) (*dto.Response, error) {
	user, err := s.findByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s - FindUserByID] error when find user by id: %w", s.name(), err)
	}

	return &dto.Response{
		Status: 200,
		Data:   &user,
	}, nil
}

func (s *UserService) FindRandomUser(ctx context.Context) (*dto.Response, error) {
	count := 0
	now := time.Now()
	start := now.AddDate(0, 0, -1)
	userId := ctx.Value("userData").(*string)

	me, err := s.findByID(ctx, *userId)
	if err != nil {
		return nil, fmt.Errorf("%s - FindRandomUser] error when find user by id: %w", s.name(), err)
	}

	val, err := s.Redis.Get(ctx, fmt.Sprintf(helper.REDIS_KEY_USER_VIEW, *userId))
	if err != nil {
		return nil, fmt.Errorf("%s - FindRandomUser] error when get count random user: %w", s.name(), err)
	}

	if val != nil {
		count, _ = strconv.Atoi(*val)
	}

	if count >= 10 && (me.Features == nil || !helper.Contains(me.Features, helper.FEATURE_NO_SWIPE_QUOTA_STRING)) {
		return &dto.Response{
			Status:  200,
			Message: "You have reached the limit of random user",
		}, nil
	}

	model, err := s.UserRepo.FindUserRandom(ctx, *userId, start, now)
	if err != nil {
		return nil, fmt.Errorf("%s - FindRandomUser] error when find random user: %w", s.name(), err)
	}

	var user dto.User
	user.Model(model)

	count += 1
	duration := redis.GetEndOfDayDuration()

	if err := s.Redis.Set(ctx, fmt.Sprintf(helper.REDIS_KEY_USER_VIEW, *userId), helper.IntToString(count), &duration); err != nil {
		return nil, fmt.Errorf("%s - FindRandomUser] error when set count random user: %w", s.name(), err)
	}

	return &dto.Response{
		Status: 200,
		Data:   &user,
	}, nil
}

func (s *UserService) Reaction(ctx context.Context, req dto.ReactionRequest) (*dto.Response, error) {
	userId := ctx.Value("userData").(*string)

	if !helper.ValidateReaction(req.Reaction) {
		return nil, fmt.Errorf("%s - Reaction] error when validate reaction", s.name())
	}

	isExist, err := s.ReactionRepo.FindByUserIdTargetIdPair(ctx, *userId, req.TargetUserID)
	if err != nil {
		return nil, fmt.Errorf("%s - Reaction] error when find reaction by user id and target user id: %w", s.name(), err)
	}

	if isExist != nil {
		if err = s.ReactionRepo.Update(ctx, req.Reaction, isExist.ID.String); err != nil {
			return nil, fmt.Errorf("%s - Reaction] error when update reaction: %w", s.name(), err)
		}
	} else {
		if _, err := s.ReactionRepo.Insert(ctx, *userId, req.TargetUserID, req.Reaction); err != nil {
			return nil, fmt.Errorf("%s - Reaction] error when insert reaction: %w", s.name(), err)
		}
	}

	return &dto.Response{
		Status:  200,
		Data:    nil,
		Message: "Success, reaction has been saved",
	}, nil
}
