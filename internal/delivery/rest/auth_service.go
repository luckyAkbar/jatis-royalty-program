package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/luckyAkbar/jatis-royalty-program/internal/model"
	"github.com/luckyAkbar/jatis-royalty-program/internal/usecase"
	"github.com/sirupsen/logrus"
)

func (s *Service) handleUserLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &model.LoginByIDAndPasswordInput{}
		if err := c.Bind(input); err != nil {
			return ErrBadRequest
		}

		session, err := s.authUsecase.LoginByIDAndPassword(c.Request().Context(), input.ID, input.Password)
		switch err {
		default:
			logrus.WithFields(logrus.Fields{
				"ctx": c.Request().Context(),
				"id":  input.ID,
			}).Error(err)

			return ErrInternal
		case usecase.ErrNotFound:
			return ErrNotFound
		case usecase.ErrUnauthorized:
			return ErrUnauthorized
		case nil:
			break
		}

		return c.JSON(http.StatusOK, session)
	}
}
