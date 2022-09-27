package model

import (
	"context"
	"time"
)

type Session struct {
	ID          int64     `gorm:"id" json:"id"`
	UserID      int64     `gorm:"user_id" json:"user_id"`
	AccessToken string    `gorm:"access_token" json:"access_token"`
	CreatedAt   time.Time `gorm:"created_at" json:"created_at"`
	ExpiredAt   time.Time `gorm:"expired_at" json:"expired_at"`
}

func (s *Session) IsAccessTokenExpired() bool {
	if s == nil {
		return true
	}

	now := time.Now()
	return now.After(s.ExpiredAt)
}

type SessionRepository interface {
	Create(ctx context.Context, session *Session) error
	FindByAccessToken(ctx context.Context, accessToken string) (*Session, error)
}
