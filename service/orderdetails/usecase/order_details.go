package usecase

import (
	"context"

	"github.com/yogabagas/jatis-BE/domain/model"
	"github.com/yogabagas/jatis-BE/domain/repository"
	"github.com/yogabagas/jatis-BE/domain/service"
)

type OrderDetailsUsecaseImpl struct {
	orderDetailsRepo repository.OrderDetaisRepository
}

type OrderDetailsUsecase interface {
	CreateOrderDetails(ctx context.Context, req []service.CreateOrderDetailsReq) (resp service.CreateOrderDetailsResp, err error)
	GetOrderDetail(context.Context) ([]service.GetOrderDetailsResp, error)
}

func NewOrderDetailsUsecase(orderDetailsRepo repository.OrderDetaisRepository) OrderDetailsUsecase {
	return &OrderDetailsUsecaseImpl{orderDetailsRepo: orderDetailsRepo}
}

func (od *OrderDetailsUsecaseImpl) CreateOrderDetails(ctx context.Context, req []service.CreateOrderDetailsReq) (resp service.CreateOrderDetailsResp, err error) {
	var reqOrderDetails []model.OrderDetail

	for _, v := range req {
		reqOrderDetail := model.OrderDetail{
			OrderId:   v.OrderId,
			ProductId: v.ProductId,
			Qty:       v.Qty,
			UnitPrice: v.UnitPrice,
			Discount:  v.Discount,
			CreatedBy: v.CreatedBy,
			UpdatedBy: v.UpdatedBy,
		}
		reqOrderDetails = append(reqOrderDetails, reqOrderDetail)
	}

	err = od.orderDetailsRepo.CreateOrderDetails(ctx, reqOrderDetails)
	if err != nil {
		return service.CreateOrderDetailsResp{
			Status: "failed",
		}, err
	}

	return service.CreateOrderDetailsResp{
		Status: "created",
	}, nil
}

func (od *OrderDetailsUsecaseImpl) GetOrderDetail(ctx context.Context) (resp []service.GetOrderDetailsResp, err error) {
	res, err := od.orderDetailsRepo.ReadOrderDetails(ctx)
	if err != nil {
		return
	}

	if len(res) > 0 {
		ordDetailsResp := service.GetOrderDetailsResp{}

		for _, v := range res {
			ordDetailsResp.CustomerName = v.CustomerName
			ordDetailsResp.EmployeeName = v.EmployeeName
			ordDetailsResp.ShippingMethod = v.ShippingMethod
			ordDetailsResp.Qty = v.Qty
			ordDetailsResp.Discount = v.Discount
			ordDetailsResp.UnitPrice = v.UnitPrice
			ordDetailsResp.SubTotalPrice = v.SubTotalPrice
			ordDetailsResp.TotalPayment = v.TotalPayment
			resp = append(resp, ordDetailsResp)
		}
	}
	return resp, nil
}
