package service

import (
	"context"
	// "fmt"

	"github.com/uchupx/dating-api/src/dto"
	"github.com/uchupx/dating-api/src/repo"
	// "github.com/uchupx/kajian-api/pkg/errors"
)

type UserService struct {
	UserRepo *repo.UserRepo
}

func (UserService) name() string {
	return "UserService"
}

func (s *UserService) FindUserByID(ctx context.Context, id string) (*dto.User, error) {
	// model, err := s.UserRepo.FindUserByID(ctx, id)
	// if err != nil {
	// 	return nil, fmt.Errorf("%s - FindUserByID] error when find user by id: %w", s.name(), err)
	// } else if model == nil {
	// 	return nil, errors.ErrNotFound
	// }
	//
	// var user dto.User
	// user.Model(model)
	//
	// return &user, nil
	return nil, nil
}

func (s *UserService) Me(ctx context.Context) (*dto.User, error) {
	return nil, nil
}
