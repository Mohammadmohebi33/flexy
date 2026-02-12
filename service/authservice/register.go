package authservice

import (
	"context"
	"flexy/dto"
	"flexy/entity"
	"fmt"
	"time"
)

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {

	user := entity.User{
		ID:       0,
		Email:    req.Email,
		Name:     req.Name,
		Password: getMD5Hash(req.Password),
	}

	createdUser, err := s.repo.RegisterUser(user)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	otp := generateOTP()

	key := fmt.Sprintf("verify:%d", createdUser.ID)

	ctx := context.Background()

	err = s.redisAdapter.Client().Set(ctx, key, otp, 3*time.Minute).Err()
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("failed to store otp: %w", err)
	}

	return dto.RegisterResponse{
		ID:    uint(createdUser.ID),
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}, nil

}
