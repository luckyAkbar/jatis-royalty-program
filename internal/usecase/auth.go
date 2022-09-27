package usecase

import (
	"context"
	"time"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/jatis-royalty-program/internal/config"
	"github.com/luckyAkbar/jatis-royalty-program/internal/helper"
	"github.com/luckyAkbar/jatis-royalty-program/internal/model"
	"github.com/luckyAkbar/jatis-royalty-program/internal/repository"
	"github.com/sirupsen/logrus"
)

type authUsecase struct {
	sessionRepo model.SessionRepository
	userRepo    model.UserRepository
}

func NewAuthUsecase(sessionRepo model.SessionRepository, userRepo model.UserRepository) model.AuthUsecase {
	return &authUsecase{
		sessionRepo,
		userRepo,
	}
}

func (u *authUsecase) LoginByIDAndPassword(ctx context.Context, id int64, password string) (*model.Session, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
		"id":  id,
	})

	user, err := u.userRepo.GetByID(ctx, id)
	switch err {
	default:
		logger.Error(err)
		return nil, ErrInternal
	case repository.ErrNotFound:
		return nil, ErrNotFound
	case nil:
		break
	}

	if !helper.IsHashedStringMatch([]byte(password), []byte(user.Password)) {
		logger.Info("invalid password usage on user: ", user.ID)
		return nil, ErrUnauthorized
	}

	session := &model.Session{
		ID:          utils.GenerateID(),
		UserID:      id,
		AccessToken: helper.GenerateToken(user.ID),
		CreatedAt:   time.Now(),
		ExpiredAt:   time.Now().Add(config.DefaultAccessTokenExpiry),
	}

	if err := u.sessionRepo.Create(ctx, session); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	return session, nil
}
