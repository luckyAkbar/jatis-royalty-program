package rest

import (
	"net/http"

	"github.com/kumparan/go-utils"
	"github.com/labstack/echo/v4"
	"github.com/luckyAkbar/jatis-royalty-program/internal/model"
	"github.com/luckyAkbar/jatis-royalty-program/internal/usecase"
	"github.com/sirupsen/logrus"
)

func (s *Service) handleRedeemVoucherByInvoice() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &model.ClaimVoucherInput{}
		if err := c.Bind(input); err != nil {
			return ErrBadRequest
		}

		voucher, err := s.voucherUsecase.ClaimByInvoiceID(c.Request().Context(), input.InvoiceID)
		switch err {
		default:
			logrus.WithFields(logrus.Fields{
				"ctx":   utils.DumpIncomingContext(c.Request().Context()),
				"input": utils.Dump(input),
			}).Error(err)
			return ErrInternal
		case usecase.ErrNotFound:
			return ErrNotFound
		case usecase.ErrPreconditionFailed:
			return ErrPreconditionFailed
		case nil:
			return c.JSON(http.StatusOK, voucher)
		}
	}
}
