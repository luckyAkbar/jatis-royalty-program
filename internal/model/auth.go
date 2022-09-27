package model

import "context"

type LoginByIDAndPasswordInput struct {
	ID int64 `json:"id"`
	Password string `json:"password"`
}

type AuthUsecase interface {
	LoginByIDAndPassword(ctx context.Context, id int64, password string) (*Session, error)
}