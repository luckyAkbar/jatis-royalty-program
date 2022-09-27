package model

import (
	"context"
	"errors"
	"time"

	"github.com/luckyAkbar/jatis-royalty-program/internal/helper"
)

type User struct {
	ID        int64     `gorm:"id" json:"id"`
	UserID    int64     `gorm:"user_id" json:"user_id"`
	Password  string    `gorm:"password" json:"-"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
}

func (u *User) Encrypt() error {
	password, err := helper.HashString(u.Password)
	if err != nil {
		return errors.New("failed to encrypt password")
	}

	u.Password = password

	return nil
}

type UserUsecase interface {
	GetByID(ctx context.Context, id int64) (*User, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int64) (*User, error)
}
