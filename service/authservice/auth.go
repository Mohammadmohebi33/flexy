package authservice

import (
	"crypto/md5"
	"encoding/hex"
	"flexy/entity"
)

type Repository interface {
	RegisterUser(u entity.User) (entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
