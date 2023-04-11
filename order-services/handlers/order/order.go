package order

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rulanugrh/alpha/order/entities/domain"
	"github.com/rulanugrh/alpha/order/entities/web"
	"github.com/rulanugrh/alpha/order/services/order"
)

type ordercontroller struct {
	orders order.OrderService
}

func NewOrderController(order order.OrderService) OrderController {
	return &ordercontroller{
		orders: order,
	}
}

func (or *ordercontroller) CreateOrder(ctx echo.Context) error {
	var order domain.Order
	ctx.Bind(&order)

	response, err := or.orders.CreateOrder(order)
	if err != nil {
		return ctx.JSON(400, web.OrderFailure{
			Code:   400,
			Status: "cant create order",
		})
	}

	return ctx.JSON(201, web.OrderSuccess{
		Code:   201,
		Status: "Success Create Order",
		Data:   response,
	})
}

func (or *ordercontroller) FindId(ctx echo.Context) error {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	response, err := or.orders.FindId(uint(id))
	if err != nil {
		return ctx.JSON(400, web.OrderFailure{
			Code:   400,
			Status: "cant find order by this id",
		})
	}

	return ctx.JSON(200, web.OrderSuccess{
		Code:   200,
		Status: "success find book",
		Data:   response,
	})
}

func (or *ordercontroller) FindAll(ctx echo.Context) error {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	response, err := or.orders.FindAll(uint(id))
	if err != nil {
		return ctx.JSON(400, web.OrderFailure{
			Code:   400,
			Status: "data not found by this user id",
		})
	}

	return ctx.JSON(200, web.OrderSuccess{
		Code:   200,
		Status: "data found",
		Data:   response,
	})
}

func (or *ordercontroller) AddCart(ctx echo.Context) error {
	var order domain.OrderItem
	ctx.Bind(&order)

	response, err := or.orders.AddCart(order)
	if err != nil {
		return ctx.JSON(400, web.OrderFailure{
			Code:   400,
			Status: "cant add to cart",
		})
	}

	return ctx.JSON(201, web.OrderSuccess{
		Code:   201,
		Status: "success add to cart",
		Data:   response,
	})
}

func (or *ordercontroller) Checkout(ctx echo.Context) error {
	var order domain.Order
	ctx.Bind(&order)

	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	response, err := or.orders.Checkout(uint(id), order)
	if err != nil {
		return ctx.JSON(400, web.OrderFailure{
			Code:   400,
			Status: "cant checkout this order, maybe id is false",
		})
	}

	return ctx.JSON(201, web.OrderSuccess{
		Code:   200,
		Status: "success checkout",
		Data:   response,
	})
}

func (or *ordercontroller) ListCart(ctx echo.Context) error {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	response, err := or.orders.ListCart(uint(id))
	if err != nil {
		return ctx.JSON(400, web.OrderFailure{
			Code:   400,
			Status: "cant find cart by this id",
		})
	}

	return ctx.JSON(200, web.OrderSuccess{
		Code:   200,
		Status: "data found",
		Data:   response,
	})
}

func (or *ordercontroller) ListNotPaid(ctx echo.Context) error {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	response, err := or.orders.ListNotPaid(uint(id))
	if err != nil {
		return ctx.JSON(400, web.OrderFailure{
			Code:   400,
			Status: "cant find by this id",
		})
	}

	return ctx.JSON(200, web.OrderSuccess{
		Code:   200,
		Status: "data found",
		Data:   response,
	})
}

func (or *ordercontroller) DeleteCart(ctx echo.Context) error {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	err := or.orders.DeleteCart(uint(id))
	if err != nil {
		return ctx.JSON(400, web.OrderFailure{
			Code:   400,
			Status: "cant delete this cart by this id",
		})
	}

	return ctx.JSON(200, web.OrderSuccess{
		Code:   200,
		Status: "success delete cart",
		Data:   nil,
	})
}
