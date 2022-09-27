package console

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luckyAkbar/jatis-royalty-program/auth"
	"github.com/luckyAkbar/jatis-royalty-program/internal/config"
	"github.com/luckyAkbar/jatis-royalty-program/internal/db"
	"github.com/luckyAkbar/jatis-royalty-program/internal/delivery/rest"
	"github.com/luckyAkbar/jatis-royalty-program/internal/repository"
	"github.com/luckyAkbar/jatis-royalty-program/internal/usecase"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var runServer = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Long:  "Use this command to start Himatro API HTTP server",
	Run:   InitServer,
}

func init() {
	RootCmd.AddCommand(runServer)
}

// InitServer initialize HTTP server
func InitServer(_ *cobra.Command, _ []string) {
	db.InitializePostgresConn()
	setupLogger()

	sqlDB, err := db.PostgresDB.DB()
	if err != nil {
		logrus.Fatal("unable to start server. reason: ", err.Error())
	}
	defer sqlDB.Close()

	sessionRepo := repository.NewSessionRepository(db.PostgresDB)
	userRepo := repository.NewUserRepository(db.PostgresDB)
	invoiceRepo := repository.NewInvoiceRepository(db.PostgresDB)
	voucherRepo := repository.NewVoucherRepository(db.PostgresDB)

	authUsecase := usecase.NewAuthUsecase(sessionRepo, userRepo)
	invoiceUsecase := usecase.NewInvoiceUsecase(invoiceRepo)
	voucherUsecase := usecase.NewVoucherUsecase(voucherRepo, invoiceRepo)

	authMiddleware := auth.NewMiddleware(sessionRepo, userRepo)

	HTTPServer := echo.New()

	HTTPServer.Pre(middleware.AddTrailingSlash())
	HTTPServer.Use(middleware.Logger())
	HTTPServer.Use(authMiddleware.UserSessionMiddleware())
	HTTPServer.Use(authMiddleware.RejectUnauthorizedRequest())

	RESTGroup := HTTPServer.Group("rest")

	rest.NewRESTService(authUsecase, invoiceUsecase, voucherUsecase, RESTGroup)

	if err := HTTPServer.Start(fmt.Sprintf(":%s", config.ServerPort())); err != nil {
		logrus.Fatal("unable to start server. reason: ", err.Error())
	}

	logrus.Info("Server running on port: ", config.ServerPort())
}
