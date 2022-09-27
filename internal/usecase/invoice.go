package usecase

import (
	"context"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/jatis-royalty-program/auth"
	"github.com/luckyAkbar/jatis-royalty-program/internal/model"
	"github.com/sirupsen/logrus"
)

type invoiceUsecase struct {
	invoiceRepo model.InvoiceRepository
}

func NewInvoiceUsecase(invoiceRepo model.InvoiceRepository) model.InvoiceUsecase {
	return &invoiceUsecase{
		invoiceRepo,
	}
}

func (u *invoiceUsecase) Create(ctx context.Context, totalPrice int64) (*model.Invoice, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":        utils.DumpIncomingContext(ctx),
		"totalPrice": totalPrice,
	})

	user := auth.GetUserFromCtx(ctx)
	invoice := &model.Invoice{
		ID:         utils.GenerateID(),
		UserID:     user.ID,
		TotalPrice: totalPrice,
	}
	if err := u.invoiceRepo.Create(ctx, invoice); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	return invoice, nil
}
