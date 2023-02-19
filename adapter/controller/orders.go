package controller

import (
	"context"

	"os"

	"github.com/gocarina/gocsv"
	"github.com/yogabagas/jatis-BE/domain/service"
	orders "github.com/yogabagas/jatis-BE/service/orders/usecase"
)

type OrdersControllerImpl struct {
	ordersUsecase orders.OrdersUsecase
}

type OrdersController interface {
	CreateOrders(ctx context.Context, req []service.CreateOrdersReq) (resp service.CreateOrdersResp, err error)
}

func NewOrdersController(ordersUsecase orders.OrdersUsecase) OrdersController {
	return &OrdersControllerImpl{ordersUsecase: ordersUsecase}
}

func (cc *OrdersControllerImpl) CreateOrders(ctx context.Context, req []service.CreateOrdersReq) (resp service.CreateOrdersResp, err error) {

	file, err := os.Open("csv_orders.csv")
	if err != nil {
		return service.CreateOrdersResp{
			Status: "failed",
		}, err
	}

	if err := gocsv.Unmarshal(file, &req); err != nil {
		return service.CreateOrdersResp{
			Status: "failed",
		}, err
	}

	return cc.ordersUsecase.CreateOrders(ctx, req)
}
