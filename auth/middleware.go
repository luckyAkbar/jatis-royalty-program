package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	models "github.com/luckyAkbar/jatis-royalty-program/internal/model"
	"github.com/sirupsen/logrus"
)

type Middleware struct {
	sessionRepo models.SessionRepository
	userRepo    models.UserRepository
}

func NewMiddleware(sessionRepo models.SessionRepository, userRepo models.UserRepository) *Middleware {
	return &Middleware{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

func (am *Middleware) UserSessionMiddleware() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := am.getTokenFromHeader(c.Request())
			if token == "" {
				return hf(c)
			}

			ctx := c.Request().Context()

			session, err := am.getSessionFromAccessToken(ctx, token)
			if err != nil {
				return hf(c)
			}

			// just pass if expired. The next middleware should block the request
			// if needed
			if session.IsAccessTokenExpired() {
				return hf(c)
			}

			user, err := am.userRepo.GetByID(ctx, session.UserID)
			if err != nil {
				return hf(c)
			}

			ctx = SetUserToCtx(ctx, User{
				ID: user.ID,
			})

			c.SetRequest(c.Request().WithContext(ctx))

			return hf(c)
		}
	}
}

func (am *Middleware) RejectUnauthorizedRequest() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			url := ctx.Request().URL
			if url.Path == "/rest/auth/login/" {
				return hf(ctx)
			}

			user := GetUserFromCtx(ctx.Request().Context())
			if user == nil {
				return ErrUnauthorized
			}

			return hf(ctx)
		}
	}
}

func (am *Middleware) getTokenFromHeader(req *http.Request) string {
	authHeader := strings.Split(req.Header.Get(_headerAuthorization), " ")

	if len(authHeader) != 2 || authHeader[0] != _authScheme {
		return ""
	}

	return strings.TrimSpace(authHeader[1])
}

func (am *Middleware) getSessionFromAccessToken(ctx context.Context, token string) (*models.Session, error) {
	logger := logrus.WithField("token", token)

	session, err := am.sessionRepo.FindByAccessToken(ctx, token)
	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return session, nil
}
