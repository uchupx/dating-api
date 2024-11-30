package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/uchupx/dating-api/pkg/database/redis"
	"github.com/uchupx/dating-api/pkg/helper"
	"github.com/uchupx/dating-api/pkg/jwt"
	"github.com/uchupx/dating-api/src/dto"
	"github.com/uchupx/dating-api/src/repo"

	"github.com/uchupx/kajian-api/pkg/errors"
)

type AuthService struct {
	UserRepo         *repo.UserRepo
	ClientRepo       *repo.ClientRepo
	RefreshTokenRepo *repo.RefreshTokenRepo
	JWT              jwt.CryptService
	Redis            *redis.Redis
}

func (AuthService) name() string {
	return "AuthService"
}

func (s *AuthService) Login(ctx context.Context, req dto.AuthRequest) (*dto.Response, error) {
	var user dto.User
	if req.Username == nil || req.Password == nil {
		return nil, fmt.Errorf("%s - Login] username or password is required", s.name())
	}

	client, err := s.ClientRepo.FindAppsByKey(ctx, req.ClientId)
	if err != nil {
		return nil, fmt.Errorf("%s - Login] error when find client app: %w", s.name(), err)
	} else if client == nil {
		return nil, errors.ErrNotFound
	}

	isClientValid, err := s.JWT.Verify(req.ClientSecret, client.Secret.String)
	if err != nil {
		return nil, fmt.Errorf("%s - Login] error when verify client secret: %w", s.name(), err)
	}

	if !isClientValid {
		return nil, errors.ErrUnauthorized
	}

	model, err := s.UserRepo.FindUserByUsernameEmail(ctx, *req.Username)
	if err != nil {
		return nil, fmt.Errorf("%s - Login] error when find user by username: %w", s.name(), err)
	} else if model == nil {
		return nil, errors.ErrNotFound
	}

	isValid, err := s.JWT.Verify(*req.Password, model.Password.String)
	if err != nil {
		return nil, fmt.Errorf("%s - Login] error when verify value: %w", s.name(), err)
	}

	if !isValid {
		return nil, errors.ErrUnauthorized
	}

	user.Model(model)

	token, err := s.JWT.CreateAccessToken(1*time.Hour, user)
	if err != nil {
		return nil, fmt.Errorf("%s - Login] error when create access token: %w", s.name(), err)
	}

	duration := 1 * time.Hour

	if err := s.Redis.Set(ctx, fmt.Sprintf("%s:%s", helper.REDIS_KEY_AUTH, *token), user.ID, &duration); err != nil {
		return nil, fmt.Errorf("%s - Login] error when set redis: %w", s.name(), err)
	}

	// refresh token

	refreshToken, err := s.JWT.CreateAccessToken(24*time.Hour, user)
	if err != nil {
		return nil, fmt.Errorf("%s - Login] error when create refresh token: %w", s.name(), err)
	}

	exp := time.Now().Add(24 * time.Hour)

	if _, err := s.RefreshTokenRepo.Insert(ctx, user.ID, user.ClientAppId, *token, exp); err != nil {
		return nil, fmt.Errorf("%s - Login] error when insert refresh token: %w", s.name(), err)
	}

	return &dto.Response{
		Status: 200,
		Data: dto.TokenResponse{
			Token:        *token,
			RefreshToken: *refreshToken,
			Expired:      int64(duration.Seconds()),
		},
	}, nil
}

func (s *AuthService) SignUp(ctx context.Context, req dto.SignUpRequest) (*dto.Response, error) {
	client, err := s.ClientRepo.FindAppsByKey(ctx, req.ClientKey)
	if err != nil {
		return nil, fmt.Errorf("%s - SignUp] error when find client app: %w", s.name(), err)
	} else if client == nil {
		return nil, errors.ErrNotFound
	}

	signPassword, err := s.JWT.CreateSignPSS(req.Password)
	if err != nil {
		return nil, fmt.Errorf("%s - SignUp] error when create signature password: %w", s.name(), err)
	}

	now := time.Now()

	newUser := dto.User{
		Username:    req.Username,
		Password:    signPassword,
		Email:       req.Email,
		ClientAppId: client.ID.String,
		Created:     now,
	}

	id, err := s.UserRepo.Insert(ctx, newUser.ToModel())
	if err != nil {
		return nil, fmt.Errorf("%s - SignUp] error when creating user: %w", s.name(), err)
	}

	newUser.ID = *id

	return &dto.Response{
		Status: 201,
		Data: dto.EntityResponse{
			Id:     *id,
			Entity: "users",
		},
	}, nil
}

func (s *AuthService) RetrieveUser(ctx context.Context, token string) (*dto.User, error) {
	resToken, err := s.JWT.VerifyJWTToken(token)
	if err != nil {
		return nil, fmt.Errorf("%s - RetrieveUser] error when verify token: %w", s.name(), err)
	}

	bytes, err := json.Marshal(resToken)
	if err != nil {
		return nil, fmt.Errorf("%s - RetrieveUser] error when marshal token: %w", s.name(), err)
	}

	var user dto.User

	if err := json.Unmarshal(bytes, &user); err != nil {
		return nil, fmt.Errorf("%s - RetrieveUser] error when unmarshal token: %w", s.name(), err)
	}

	return &user, nil
}

func (s *AuthService) AddClient(ctx context.Context, req dto.ClientPost) (*dto.Response, error) {

	secret := RandomString(20)
	clientSecret, err := s.JWT.CreateSignPSS(secret)
	if err != nil {
		return nil, fmt.Errorf("%s - AddClient] error when create signature password: %w", s.name(), err)
	}

	data := dto.Client{
		Name:   req.Name,
		Key:    RandomString(20),
		Secret: clientSecret,
	}

	id, err := s.ClientRepo.Insert(ctx, data.ToModel())
	if err != nil {
		return nil, fmt.Errorf("%s - AddClient] error when insert client: %w", s.name(), err)
	}

	return &dto.Response{
		Status: 201,
		Data: dto.EntityResponse{
			Id:     id,
			Entity: "client_apps",
			Meta: map[string]interface{}{
				"secret": secret,
				"key":    data.Key,
			},
		},
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (*dto.Response, error) {
	client, err := s.ClientRepo.FindAppsByKey(ctx, req.ClientId)
	if err != nil {
		return nil, fmt.Errorf("%s - RefreshToken] error when find client app: %w", s.name(), err)
	} else if client == nil {
		return nil, errors.ErrNotFound
	}

	isClientValid, err := s.JWT.Verify(req.ClientSecret, client.Secret.String)
	if err != nil {
		return nil, fmt.Errorf("%s - RefreshToken] error when verify client secret: %w", s.name(), err)
	}
	if !isClientValid {
		return nil, errors.ErrUnauthorized
	}

	user, err := s.RetrieveUser(ctx, req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("%s - RefreshToken] error when retrieve user: %w", s.name(), err)
	}

	token, err := s.JWT.CreateAccessToken(1*time.Hour, *user)
	if err != nil {
		return nil, fmt.Errorf("%s - RefreshToken] error when create access token: %w", s.name(), err)
	}

	duration := 1 * time.Hour
	if err := s.Redis.Set(ctx, fmt.Sprintf("%s:%s", helper.REDIS_KEY_AUTH, *token), user.ID, &duration); err != nil {
		return nil, fmt.Errorf("%s - RefreshToken] error when set redis: %w", s.name(), err)
	}

	return &dto.Response{
		Status: 200,
		Data: dto.TokenResponse{
			Token:   *token,
			Expired: int64(duration.Seconds()),
		},
	}, nil
}

func RandomString(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
