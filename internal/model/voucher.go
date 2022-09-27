package model

import (
	"context"
	"time"
)

type Voucher struct {
	ID        int64     `gorm:"id" json:"id"`
	UserID    int64     `gorm:"user_id" json:"user_id"`
	InvoiceID int64     `gorm:"invoice_id" json:"invoice_id"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	ExpiredAt time.Time `gorm:"expired_at" json:"expired_at"`
	IsClaimed bool      `gorm:"is_claimed" json:"is_claimed"`
	Value     int64     `gorm:"value" json:"value"`
}

type ClaimVoucherInput struct {
	InvoiceID int64 `json:"invoice_id"`
}

type VoucherRepository interface {
	GetByInvoiceID(ctx context.Context, id int64) (*Voucher, error)
	Create(ctx context.Context, voucher *Voucher) error
}

type VoucherUsecase interface {
	ClaimByInvoiceID(ctx context.Context, invoiceID int64) (*Voucher, error)
}
