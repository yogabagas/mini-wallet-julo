package usecase

import (
	"context"

	"github.com/yogabagas/jatis-BE/domain/model"
	"github.com/yogabagas/jatis-BE/domain/repository"
	"github.com/yogabagas/jatis-BE/domain/service"
)

type ProductsUsecaseImpl struct {
	productsRepo repository.ProductsRepository
}

type ProductsUsecase interface {
	CreateProducts(ctx context.Context, req []service.CreateProductsReq) (resp service.CreateProductsResp, err error)
}

func NewProductsUsecase(productsRepo repository.ProductsRepository) ProductsUsecase {
	return &ProductsUsecaseImpl{productsRepo: productsRepo}
}

func (pu *ProductsUsecaseImpl) CreateProducts(ctx context.Context, req []service.CreateProductsReq) (resp service.CreateProductsResp, err error) {
	var reqProducts []model.Product

	for _, v := range req {
		reqProduct := model.Product{
			ProductName: v.ProductName,
			UnitPrice:   v.UnitPrice,
			InStock:     v.InStock,
			CreatedBy:   v.CreatedBy,
			UpdatedBy:   v.UpdatedBy,
		}
		reqProducts = append(reqProducts, reqProduct)
	}

	err = pu.productsRepo.InsertProducts(ctx, reqProducts)
	if err != nil {
		return service.CreateProductsResp{
			Status: "failed",
		}, err
	}

	return service.CreateProductsResp{
		Status: "created",
	}, nil
}
