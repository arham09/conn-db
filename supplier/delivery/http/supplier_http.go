package http

import (
	"context"
	"net/http"

	"github.com/arham09/conn-db/helpers"
	"github.com/arham09/conn-db/middleware"
	"github.com/arham09/conn-db/supplier"
	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type SupplierHandler struct {
	SUsecase supplier.Usecase
}

func NewSupplierHandler(e *echo.Echo, us supplier.Usecase, middleware *middleware.GoMiddleware) {
	handler := &SupplierHandler{
		SUsecase: us,
	}
	e.GET("/supplier", handler.FetchAll, middleware.CORS)
}

func (s *SupplierHandler) FetchAll(c echo.Context) error {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	listSup, err := s.SUsecase.FetchAll(ctx)

	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listSup)
}
