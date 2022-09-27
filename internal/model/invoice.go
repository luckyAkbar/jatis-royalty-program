package model

import (
	"context"

	"github.com/luckyAkbar/jatis-royalty-program/internal/config"
)

type Invoice struct {
	ID         int64 `gorm:"id" json:"id"`
	UserID     int64 `gorm:"user_id" json:"user_id"`
	TotalPrice int64 `gorm:"total" json:"total_price"`
}

func (i *Invoice) IsWorthForVoucherAward() bool {
	return i.TotalPrice > config.VoucherAwardThreshold()
}

type CreateInvoiceInput struct {
	TotalPrice int64 `json:"total_price"`
}

type InvoiceRepository interface {
	GetByID(ctx context.Context, id int64) (*Invoice, error)
	Create(ctx context.Context, invoice *Invoice) error
}

type InvoiceUsecase interface {
	Create(ctx context.Context, totalPrice int64) (*Invoice, error)
}
