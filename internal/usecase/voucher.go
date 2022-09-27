package usecase

import (
	"context"
	"time"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/jatis-royalty-program/auth"
	"github.com/luckyAkbar/jatis-royalty-program/internal/config"
	"github.com/luckyAkbar/jatis-royalty-program/internal/model"
	"github.com/luckyAkbar/jatis-royalty-program/internal/repository"
	"github.com/sirupsen/logrus"
)

type voucherUsecase struct {
	voucherRepo model.VoucherRepository
	invoiceRepo model.InvoiceRepository
}

func NewVoucherUsecase(voucherRepo model.VoucherRepository, invoiceRepo model.InvoiceRepository) model.VoucherUsecase {
	return &voucherUsecase{
		voucherRepo,
		invoiceRepo,
	}
}

func (u *voucherUsecase) ClaimByInvoiceID(ctx context.Context, invoiceID int64) (*model.Voucher, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":       utils.DumpIncomingContext(ctx),
		"invoiceID": invoiceID,
	})

	invoice, err := u.invoiceRepo.GetByID(ctx, invoiceID)
	switch err {
	default:
		logger.Error(err)
		return nil, ErrInternal
	case repository.ErrNotFound:
		return nil, ErrNotFound
	case nil:
		break
	}

	user := auth.GetUserFromCtx(ctx)
	if user.ID != invoice.UserID {
		return nil, ErrUnauthorized
	}

	if !invoice.IsWorthForVoucherAward() {
		return nil, ErrPreconditionFailed
	}

	voucher := &model.Voucher{
		ID:        utils.GenerateID(),
		UserID:    invoice.UserID,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(config.DefaultVoucherExpiry),
		IsClaimed: false,
		Value:     config.VouherAwardPriceValue(),
		InvoiceID: invoice.ID,
	}

	if err := u.voucherRepo.Create(ctx, voucher); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	return voucher, nil
}
