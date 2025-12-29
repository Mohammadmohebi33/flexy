package entity

import (
	"time"
)

type User struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Password        string     `json:"-"`
	ScoreWeek       int        `json:"score_week"`
	ScoreMonth      int        `json:"score_month"`
	Score           int        `json:"score"`
	AvatarURL       string     `json:"avatar_url"`
	Status          string     `json:"status"`
	ActiveDays      int        `json:"active_days"`
	LastActiveAt    *time.Time `json:"last_active_at"`
	RememberToken   string     `json:"remember_token"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
