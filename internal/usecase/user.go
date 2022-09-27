package usecase

import (
	"context"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/jatis-royalty-program/internal/model"
	"github.com/luckyAkbar/jatis-royalty-program/internal/repository"
	"github.com/sirupsen/logrus"
)

type userUsecase struct {
	userRepo model.UserRepository
}

func NewUserUsecase(userRepo model.UserRepository) model.UserUsecase {
	return &userUsecase{
		userRepo,
	}
}

func (u *userUsecase) GetByID(ctx context.Context, id int64) (*model.User, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
		"id": id,
	})

	user, err := u.userRepo.GetByID(ctx, id)
	switch err {
	default:
		logger.Error(err)
		return nil, ErrInternal
	case repository.ErrNotFound:
		return nil, ErrNotFound
	case nil:
		return user, nil
	}
}