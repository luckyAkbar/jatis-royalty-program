package repository

import (
	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/jatis-royalty-program/internal/model"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type voucherRepo struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) model.VoucherRepository {
	return &voucherRepo{
		db,
	}
}

func (r *voucherRepo) Create(ctx context.Context, voucher *model.Voucher) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":    utils.DumpIncomingContext(ctx),
		"voucer": utils.Dump(voucher),
	})

	if err := r.db.WithContext(ctx).Create(voucher).Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (r *voucherRepo) GetByInvoiceID(ctx context.Context, id int64) (*model.Voucher, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
		"id":  id,
	})

	voucher := &model.Voucher{}
	err := r.db.WithContext(ctx).Model(&model.Voucher{}).Where("invoice_id = ?", id).Take(voucher).Error
	switch err {
	default:
		logger.Error(err)
		return nil, err
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	case nil:
		return voucher, nil
	}
}
