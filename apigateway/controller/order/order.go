package order

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rulanugrh/alpha/apigateway/entity/order"
	"github.com/rulanugrh/alpha/apigateway/entity/web"
	orderServ "github.com/rulanugrh/alpha/apigateway/services/order"
)

type ordercontroller struct {
	order orderServ.OrderInterfaces
}

func NewOrderContoller(order orderServ.OrderInterfaces) OrderInterfaces {
	return &ordercontroller{
		order: order,
	}
}

func (or *ordercontroller) Create(ctx echo.Context) error {
	var order order.OrderModel
	ctx.Bind(&order)

	resp, err := or.order.Create(order)
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "cannot create data",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Get Data",
		Data:   resp,
	}

	return ctx.JSON(200, response)
}

func (or *ordercontroller) FindId(ctx echo.Context) error {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	resp, err := or.order.FindId(uint(id))
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "cannot create data",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Get Data",
		Data:   resp,
	}

	return ctx.JSON(200, response)
}

func (or *ordercontroller) FindAll(ctx echo.Context) error {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	resp, err := or.order.FindAll(uint(id))
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "cannot create data",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Get Data",
		Data:   resp,
	}

	return ctx.JSON(200, response)
}

func (or *ordercontroller) AddCart(ctx echo.Context) error {
	var order order.OrderItemModel
	ctx.Bind(&order)

	resp, err := or.order.AddCart(order)
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "cannot create data",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Get Data",
		Data:   resp,
	}

	return ctx.JSON(200, response)
}

func (or *ordercontroller) Checkout(ctx echo.Context) error {
	var order order.OrderModel
	ctx.Bind(&order)

	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	resp, err := or.order.Checkout(uint(id), order)
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "cannot create data",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Get Data",
		Data:   resp,
	}

	return ctx.JSON(200, response)
}

func (or *ordercontroller) ListCart(ctx echo.Context) error {
	getId := ctx.Param("userId")
	id, _ := strconv.Atoi(getId)

	resp, err := or.order.ListCart(uint(id))
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "cannot create data",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Get Data",
		Data:   resp,
	}

	return ctx.JSON(200, response)
}

func (or *ordercontroller) DeleteCart(ctx echo.Context) error {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	resp, err := or.order.DeleteCart(uint(id))
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "cannot create data",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Get Data",
		Data:   resp,
	}

	return ctx.JSON(200, response)
}

func (or *ordercontroller) ListNotPaid(ctx echo.Context) error {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	resp, err := or.order.ListNotPaid(uint(id))
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "cannot create data",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Get Data",
		Data:   resp,
	}

	return ctx.JSON(200, response)
}
