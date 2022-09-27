package rest

import (
	"net/http"

	"github.com/kumparan/go-utils"
	"github.com/labstack/echo/v4"
	"github.com/luckyAkbar/jatis-royalty-program/internal/model"
	"github.com/sirupsen/logrus"
)

func (s *Service) handleCreateInvoice() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &model.CreateInvoiceInput{}
		if err := c.Bind(input); err != nil {
			return ErrBadRequest
		}

		invoice, err := s.invoiceUsecase.Create(c.Request().Context(), input.TotalPrice)
		switch err {
		default:
			logrus.WithFields(logrus.Fields{
				"ctx":   utils.DumpIncomingContext(c.Request().Context()),
				"input": utils.Dump(input),
			}).Error(err)
			return ErrInternal
		case nil:
			return c.JSON(http.StatusOK, invoice)
		}
	}
}
