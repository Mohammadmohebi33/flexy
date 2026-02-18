package userservice

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"flexy/adapter/redis"
	"flexy/dto"
	"flexy/entity"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	RegisterUser(u entity.User) (entity.User, error)
	LoginUser(u entity.User) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	auth         AuthGenerator
	repo         Repository
	redisAdapter redis.Adapter
}

func New(authGenerator AuthGenerator, repo Repository) Service {
	return Service{
		auth: authGenerator,
		repo: repo,
		redisAdapter: redis.New(
			redis.Config{
				Host:     "127.0.0.1",
				Port:     6379,
				Password: "",
				DB:       0,
			}),
	}
}

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user := entity.User{
		ID:       0,
		Email:    req.Email,
		Name:     req.Name,
		Password: hashedPassword,
	}

	createdUser, err := s.repo.RegisterUser(user)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("failed to register user: %w", err)
	}

	otp, err := generateOTP()
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("failed to generate otp: %w", err)
	}

	key := fmt.Sprintf("verify:%d", createdUser.ID)
	ctx := context.Background()

	err = s.redisAdapter.Client().Set(ctx, key, otp, 3*time.Minute).Err()
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("failed to store otp: %w", err)
	}

	// TODO: send OTP to user's email here (e.g. via email service)

	return dto.RegisterResponse{
		ID:    uint(createdUser.ID),
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}, nil
}

func (s Service) Login(req dto.LoginRequest) (dto.LoginResponse, error) {

	user := entity.User{
		Email:    req.Email,
		Password: req.Password,
	}

	loggedInUser, err := s.repo.LoginUser(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("invalid credentials: %w", err)
	}

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return dto.LoginResponse{
		ID:    uint(loggedInUser.ID),
		Name:  loggedInUser.Name,
		Email: loggedInUser.Email,
		Tokens: dto.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func generateOTP() (string, error) {
	max := big.NewInt(1_000_000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", n.Int64()), nil
}
