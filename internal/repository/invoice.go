package repository

import (
	"context"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/jatis-royalty-program/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type invoiceRepo struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) model.InvoiceRepository {
	return &invoiceRepo{
		db,
	}
}

func (r *invoiceRepo) Create(ctx context.Context, invoice *model.Invoice) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":     utils.DumpIncomingContext(ctx),
		"invoice": utils.Dump(invoice),
	})

	if err := r.db.WithContext(ctx).Create(invoice).Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (r *invoiceRepo) GetByID(ctx context.Context, id int64) (*model.Invoice, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
		"id":  id,
	})

	invoice := &model.Invoice{}
	err := r.db.WithContext(ctx).Model(&model.Invoice{}).Where("id = ?", id).Take(invoice).Error
	switch err {
	default:
		logger.Error(err)
		return nil, err
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	case nil:
		return invoice, nil
	}
}
