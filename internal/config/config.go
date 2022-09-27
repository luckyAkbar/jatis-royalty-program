package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	DefaultTokenLength       int = 25
	DefaultAccessTokenExpiry     = time.Hour * 24 * 30
	DefaultVoucherExpiry         = time.Hour * 24 * 30 * 3
)

func ServerPort() string {
	return os.Getenv("SERVER_PORT")
}

func LogLevel() string {
	return os.Getenv("LOG_LEVEL")
}

func Env() string {
	cfg := os.Getenv("ENV")

	if cfg == "" {
		return "development"
	}

	return cfg
}

func PostgresDSN() string {
	host := os.Getenv("PG_HOST")
	db := os.Getenv("PG_DATABASE")
	user := os.Getenv("PG_USER")
	pw := os.Getenv("PG_PASSWORD")
	port := os.Getenv("PG_PORT")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pw, db, port)
}

func VoucherAwardThreshold() int64 {
	cfg := os.Getenv("VOUCHER_AWARD_PRICE_THRESHOLD")
	if cfg == "" {
		logrus.Fatal("voucher award threshold is not specified")
	}

	threshold, err := strconv.ParseInt(cfg, 10, 64)
	if err != nil {
		logrus.Fatal("voucher award threshold is not a number")
	}

	return threshold
}

func VouherAwardPriceValue() int64 {
	cfg := os.Getenv("VOUCHER_AWARD_PRICE_VALUE")
	if cfg == "" {
		logrus.Fatal("voucher award value is not specified")
	}

	value, err := strconv.ParseInt(cfg, 10, 64)
	if err != nil {
		logrus.Fatal("voucher award value is not a number")
	}

	return value
}
