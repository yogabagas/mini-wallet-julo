package controller

import (
	"context"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/yogabagas/jatis-BE/domain/service"

	orderDetails "github.com/yogabagas/jatis-BE/service/orderdetails/usecase"
)

type OrderDetailsControllerImpl struct {
	orderDetailsUsecase orderDetails.OrderDetailsUsecase
}

type OrderDetailsController interface {
	CreateOrderDetails(ctx context.Context, req []service.CreateOrderDetailsReq) (resp service.CreateOrderDetailsResp, err error)
	GetOrderDetail(context.Context) ([]service.GetOrderDetailsResp, error)
}

func NewOrderDetailsController(orderDetailsUsecase orderDetails.OrderDetailsUsecase) OrderDetailsController {
	return &OrderDetailsControllerImpl{orderDetailsUsecase: orderDetailsUsecase}
}

func (oc *OrderDetailsControllerImpl) CreateOrderDetails(ctx context.Context, req []service.CreateOrderDetailsReq) (resp service.CreateOrderDetailsResp, err error) {
	file, err := os.Open("csv_order_details.csv")
	if err != nil {
		return service.CreateOrderDetailsResp{
			Status: "failed",
		}, err
	}

	if err := gocsv.Unmarshal(file, &req); err != nil {
		return service.CreateOrderDetailsResp{
			Status: "failed",
		}, err
	}

	return oc.orderDetailsUsecase.CreateOrderDetails(ctx, req)
}

func (oc *OrderDetailsControllerImpl) GetOrderDetail(ctx context.Context) (resp []service.GetOrderDetailsResp, err error) {
	return oc.orderDetailsUsecase.GetOrderDetail(ctx)
}
