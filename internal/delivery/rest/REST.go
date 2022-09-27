package rest

import (
	echo "github.com/labstack/echo/v4"
	"github.com/luckyAkbar/jatis-royalty-program/internal/model"
)

type Service struct {
	authUsecase    model.AuthUsecase
	invoiceUsecase model.InvoiceUsecase
	voucherUsecase model.VoucherUsecase

	group *echo.Group
}

func NewRESTService(authUsecase model.AuthUsecase, invoiceUsecase model.InvoiceUsecase, voucherUsecase model.VoucherUsecase, group *echo.Group) {
	service := &Service{
		authUsecase,
		invoiceUsecase,
		voucherUsecase,
		group,
	}

	service.initRoutes()
}

func (s *Service) initRoutes() {
	s.initUserService()
	s.initVoucherService()
	s.initInvoiceService()
}

func (s *Service) initUserService() {
	s.group.POST("/auth/login/", s.handleUserLogin())
}

func (s *Service) initVoucherService() {
	s.group.POST("/vouchers/redeem/", s.handleRedeemVoucherByInvoice())
}

func (s *Service) initInvoiceService() {
	s.group.POST("/invoices/create/", s.handleCreateInvoice())
}
