package authservice

import (
	"flexy/dto"
	"flexy/entity"
	"fmt"
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

	return dto.RegisterResponse{
		ID:    uint(createdUser.ID),
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}, nil

}
