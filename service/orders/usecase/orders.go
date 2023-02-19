package usecase

import (
	"context"

	"github.com/yogabagas/jatis-BE/domain/model"
	"github.com/yogabagas/jatis-BE/domain/repository"
	"github.com/yogabagas/jatis-BE/domain/service"
)

type OrdersUsecaseImpl struct {
	ordersRepo repository.OrdersRepository
}

type OrdersUsecase interface {
	CreateOrders(ctx context.Context, req []service.CreateOrdersReq) (resp service.CreateOrdersResp, err error)
}

func NewOrdersUsecase(ordersRepo repository.OrdersRepository) OrdersUsecase {
	return &OrdersUsecaseImpl{ordersRepo: ordersRepo}
}

func (pu *OrdersUsecaseImpl) CreateOrders(ctx context.Context, req []service.CreateOrdersReq) (resp service.CreateOrdersResp, err error) {
	var reqOrders []model.Order

	for _, v := range req {
		reqOrder := model.Order{
			CustomerId:          v.CustomerId,
			EmployeeId:          v.EmployeeId,
			OrderDate:           v.OrderDate,
			PurchaseOrderNumber: v.PurchaseOrderNumber,
			ShipDate:            v.ShipDate,
			ShippingMethodId:    v.ShippingMethodId,
			FreightCharge:       v.FreightCharge,
			Taxes:               v.Taxes,
			PaymentReceived:     v.PaymentReceived,
			Comment:             v.Comment,
			CreatedBy:           v.CreatedBy,
			UpdatedBy:           v.UpdatedBy,
		}
		reqOrders = append(reqOrders, reqOrder)
	}

	err = pu.ordersRepo.InsertOrders(ctx, reqOrders)
	if err != nil {
		return service.CreateOrdersResp{
			Status: "failed",
		}, err
	}

	return service.CreateOrdersResp{
		Status: "created",
	}, nil
}
