package authservice

import (
	"crypto/md5"
	"encoding/hex"
	"flexy/adapter/redis"
	"flexy/entity"
	"fmt"
	"math/rand"
	"time"
)

type Repository interface {
	RegisterUser(u entity.User) (entity.User, error)
}

type Service struct {
	repo         Repository
	redisAdapter redis.Adapter
}

func New(repo Repository) Service {
	return Service{
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

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
