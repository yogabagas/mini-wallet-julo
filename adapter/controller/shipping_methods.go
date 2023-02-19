package controller

import (
	"context"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/yogabagas/jatis-BE/domain/service"
	shippingMethods "github.com/yogabagas/jatis-BE/service/shippingmethods/usecase"
)

type ShippingMethodsControllerImpl struct {
	shippingMethodsUsecase shippingMethods.ShippingMethodsUsecase
}

type ShippingMethodsController interface {
	CreateShippingMethods(ctx context.Context, req []service.CreateShippingMethodsReq) (resp service.CreateShippingMethodsResp, err error)
}

func NewShippingMethodsController(shippingMethodsUsecase shippingMethods.ShippingMethodsUsecase) ShippingMethodsController {
	return &ShippingMethodsControllerImpl{shippingMethodsUsecase: shippingMethodsUsecase}
}

func (pc *ShippingMethodsControllerImpl) CreateShippingMethods(ctx context.Context, req []service.CreateShippingMethodsReq) (resp service.CreateShippingMethodsResp, err error) {
	file, err := os.Open("csv_shipping_methods.csv")
	if err != nil {
		return service.CreateShippingMethodsResp{
			Status: "failed",
		}, err
	}

	if err := gocsv.Unmarshal(file, &req); err != nil {
		return service.CreateShippingMethodsResp{
			Status: "failed",
		}, err
	}

	return pc.shippingMethodsUsecase.CreateShippingMethods(ctx, req)
}
