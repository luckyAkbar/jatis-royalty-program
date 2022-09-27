package repository

import (
	"context"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/jatis-royalty-program/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &userRepository{
		db,
	}
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
		"id":  id,
	})

	user := &model.User{}
	err := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Take(user).Error
	switch err {
	default:
		logger.Error(err)
		return nil, err
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	case nil:
		break
	}

	return user, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":  utils.DumpIncomingContext(ctx),
		"user": utils.Dump(user),
	})

	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
