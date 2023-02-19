package controller

import (
	"context"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/yogabagas/jatis-BE/domain/service"
	products "github.com/yogabagas/jatis-BE/service/products/usecase"
)

type ProductsControllerImpl struct {
	productsUsecase products.ProductsUsecase
}

type ProductsController interface {
	CreateProducts(ctx context.Context, req []service.CreateProductsReq) (resp service.CreateProductsResp, err error)
}

func NewProductsController(productsUsecase products.ProductsUsecase) ProductsController {
	return &ProductsControllerImpl{productsUsecase: productsUsecase}
}

func (pc *ProductsControllerImpl) CreateProducts(ctx context.Context, req []service.CreateProductsReq) (resp service.CreateProductsResp, err error) {
	file, err := os.Open("csv_products.csv")
	if err != nil {
		return service.CreateProductsResp{
			Status: "failed",
		}, err
	}

	if err := gocsv.Unmarshal(file, &req); err != nil {
		return service.CreateProductsResp{
			Status: "failed",
		}, err
	}

	return pc.productsUsecase.CreateProducts(ctx, req)
}
