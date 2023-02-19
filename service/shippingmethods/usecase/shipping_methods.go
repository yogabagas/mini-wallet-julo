package usecase

import (
	"context"

	"github.com/yogabagas/jatis-BE/domain/model"
	"github.com/yogabagas/jatis-BE/domain/repository"
	"github.com/yogabagas/jatis-BE/domain/service"
)

type ShippingMethodsUsecaseImpl struct {
	shippingMethodsRepo repository.ShippingMethodsRepository
}

type ShippingMethodsUsecase interface {
	CreateShippingMethods(ctx context.Context, req []service.CreateShippingMethodsReq) (resp service.CreateShippingMethodsResp, err error)
}

func NewShippingMethodsUsecase(shippingMethodsRepo repository.ShippingMethodsRepository) ShippingMethodsUsecase {
	return &ShippingMethodsUsecaseImpl{shippingMethodsRepo: shippingMethodsRepo}
}

func (pu *ShippingMethodsUsecaseImpl) CreateShippingMethods(ctx context.Context, req []service.CreateShippingMethodsReq) (resp service.CreateShippingMethodsResp, err error) {
	var reqShippingMethods []model.ShippingMethod

	for _, v := range req {
		reqShippingMethod := model.ShippingMethod{
			ShippingMethodName: v.ShippingMethodName,
			CreatedBy:          v.CreatedBy,
			UpdatedBy:          v.UpdatedBy,
		}
		reqShippingMethods = append(reqShippingMethods, reqShippingMethod)
	}

	err = pu.shippingMethodsRepo.InsertShippingMethods(ctx, reqShippingMethods)
	if err != nil {
		return service.CreateShippingMethodsResp{
			Status: "failed",
		}, err
	}

	return service.CreateShippingMethodsResp{
		Status: "created",
	}, nil
}
