package http

import (
	"context"
	"net/http"
	"strconv"

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
	e.GET("/v1/supplier", handler.FetchAll, middleware.CORS, middleware.UserLimiter)
	e.GET("/v1/supplier/:id", handler.FetchById, middleware.CORS)
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

func (s *SupplierHandler) FetchById(c echo.Context) error {
	idSup, err := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ErrNotFound.Error())
	}

	id := int64(idSup)

	if ctx == nil {
		ctx = context.Background()
	}

	supp, err := s.SUsecase.FetchById(ctx, id)

	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, supp)
}
