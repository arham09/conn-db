package http

import (
	"context"
	"net/http"

	"github.com/arham09/conn-db/models"
	"github.com/arham09/conn-db/supplier"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

type SupplierHandler struct {
	SUsecase supplier.Usecase
}

func NewSupplierHandler(e *echo.Echo, us supplier.Usecase) {
	handler := &SupplierHandler{
		SUsecase: us,
	}
	e.GET("/supplier", handler.FetchAll)
}

func (s *SupplierHandler) FetchAll(c echo.Context) error {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	listSup, err := s.SUsecase.FetchAll(ctx)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listSup)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
